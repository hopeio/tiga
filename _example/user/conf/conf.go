package conf

import (
	"database/sql"
	"github.com/hopeio/lemon/initialize/gormdb/postgres"
	"github.com/hopeio/lemon/initialize/log"
	"github.com/hopeio/lemon/initialize/mail"
	"github.com/hopeio/lemon/initialize/pebble"
	"github.com/hopeio/lemon/initialize/redis"
	"github.com/hopeio/lemon/initialize/ristretto"
	"github.com/hopeio/lemon/initialize/server"

	"github.com/hopeio/lemon/utils/io/fs"
	"github.com/spf13/viper"
	"runtime"
	"time"
)

var (
	Conf      = &config{}
	Dao  *dao = &dao{}
)

type config struct {
	//自定义的配置
	Customize serverConfig
	Server    server.ServerConfig
	Log       log.LogConfig
	Viper     *viper.Viper
}

type serverConfig struct {
	Volume fs.Dir

	PassSalt string
	// 天数
	TokenMaxAge time.Duration
	TokenSecret []byte
	PageSize    int8

	LuosimaoSuperPW   string
	LuosimaoVerifyURL string
	LuosimaoAPIKey    string

	QrCodeSaveDir fs.Dir //二维码保存路径
	PrefixUrl     string
	FontSaveDir   fs.Dir //字体保存路径

}

func (c *config) Init() {
	if runtime.GOOS == "windows" {
	}

	c.Customize.TokenMaxAge = time.Second * 60 * 60 * 24 * c.Customize.TokenMaxAge
}

// dao dao.
type dao struct {
	// GORMDB 数据库连接
	GORMDB   postgres.DB
	StdDB    *sql.DB
	PebbleDB pebble.DB
	// RedisPool Redis连接池
	Redis redis.Redis
	Cache ristretto.Cache
	//elastic
	Mail mail.Mail `init:"config:mail"`
}

func (d *dao) Init() {
	db := d.GORMDB
	db.Callback().Create().Remove("gorm:save_before_associations")
	db.Callback().Create().Remove("gorm:save_after_associations")
	db.Callback().Update().Remove("gorm:save_before_associations")
	db.Callback().Update().Remove("gorm:save_after_associations")

	d.StdDB, _ = db.DB.DB()
}
