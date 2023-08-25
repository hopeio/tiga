package local

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/hopeio/lemon/utils/configor/local"
	"github.com/hopeio/lemon/utils/log"
	"os"
	"time"
)

type ReloadType string

const (
	ReloadTypeFsNotify = "fsnotify"
	ReloadTypeTimer    = "timer"
)

type Local struct {
	local.Config
	ConfigPath string
	ReloadType ReloadType `json:"reloadType" enum:"fsnotify,timer"` // 本地分为Watch和AutoReload，Watch采用系统调用通知，AutoReload定时器去查文件是否变更
}

// 本地配置
func (cc *Local) HandleConfig(handle func([]byte)) error {

	_, err := os.Stat(cc.ConfigPath)
	if os.IsNotExist(err) {
		return fmt.Errorf("找不到配置: %v", err)
	}

	var watch bool
	if cc.AutoReload && cc.ReloadType == ReloadTypeFsNotify {
		// 关掉timer
		cc.AutoReload = false
		// 启用watch
		watch = true
	}
	err = local.New(&cc.Config).Handle(handle, cc.ConfigPath)
	if err != nil {
		return fmt.Errorf("配置错误: %v", err)
	}
	if watch {
		go cc.watch(cc.ConfigPath, handle)
	}

	return nil
}

func (cc *Local) watch(adCongPath string, handle func([]byte)) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	err = watcher.Add(adCongPath)
	if err != nil {
		log.Fatal(err)
	}
	interval := make(map[string]time.Time)

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			now := time.Now()
			if now.Sub(interval[event.Name]) < time.Second {
				continue
			}
			interval[event.Name] = now
			//log.Info("event:", event)
			if event.Op&fsnotify.Write == fsnotify.Write {
				err = local.New(&cc.Config).
					Handle(handle, adCongPath)
				if err != nil {
					log.Errorf("配置错误: %v", err)
				}
				log.Info("modified file:", event.Name)
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Error("error:", err)
		}
	}
}
