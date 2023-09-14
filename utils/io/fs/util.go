package fs

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/hopeio/lemon/utils/crypto"
	"github.com/hopeio/lemon/utils/log"
	"io"
	"os"
	stdpath "path"
	"strings"
)

func Exist(filepath string) bool {
	_, err := os.Stat(filepath)
	return !os.IsNotExist(err)
}

func NotExist(filepath string) bool {
	_, err := os.Stat(filepath)
	return os.IsNotExist(err)
}

func Md5(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	hash := md5.New()
	_, err = io.Copy(hash, file)
	if err != nil {
		file.Close()
		return "", err
	}
	file.Close()
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func Md5Equal(path1, path2 string) (bool, error) {
	md51, err := Md5(path1)
	if err != nil {
		return false, err
	}
	md52, err := Md5(path2)
	if err != nil {
		return false, err
	}
	return md51 == md52, nil
}

func GetMd5Name(name string) string {
	ext := stdpath.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = crypto.EncodeMD5String(fileName)
	return fileName + ext
}

type duplicateFile struct {
	path string
	md5  string
}

// 去除目录中重复的文件,默认保留参数靠前目录中的文件
func DirDeDuplicate(dirs ...string) error {
	fileSizeMap := make(map[int64][]*duplicateFile)
	for _, tmpDir := range dirs {
		err := RangeFile(tmpDir, func(dir string, entry os.DirEntry) error {
			info, _ := entry.Info()
			path := dir + PathSeparator + entry.Name()
			duplicateFiles, ok := fileSizeMap[info.Size()]
			var entryMd5 string
			if ok {
				var err error
				entryMd5, err = Md5(path)
				if err != nil {
					return err
				}
				for _, file := range duplicateFiles {
					if file.md5 == "" {
						file.md5, err = Md5(file.path)
						if err != nil {
							return err
						}
					}
					if file.md5 == entryMd5 {
						log.Debugf("exists: %s,remove:%s", file.path, path)
						err = os.Remove(path)
						if err != nil {
							return err
						}
						return nil
					}
				}
			}
			fileSizeMap[info.Size()] = append(fileSizeMap[info.Size()], &duplicateFile{path: path, md5: entryMd5})
			return nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// 去除两个目录中重复的文件,默认保留第一个目录中的文件
func TwoDirDeDuplicate(dir1, dir2 string) error {
	fileSizeMap := make(map[int64][]*duplicateFile)
	err := RangeFile(dir1, func(dir string, entry os.DirEntry) error {
		info, _ := entry.Info()
		fileSizeMap[info.Size()] = append(fileSizeMap[info.Size()], &duplicateFile{path: dir + PathSeparator + entry.Name()})
		return nil
	})

	if err != nil {
		return err
	}

	return RangeFile(dir2, func(dir string, entry os.DirEntry) error {
		info, _ := entry.Info()
		if duplicateFiles, ok := fileSizeMap[info.Size()]; ok {
			path := dir + PathSeparator + entry.Name()
			entryMd5, err := Md5(path)
			if err != nil {
				return err
			}
			for _, file := range duplicateFiles {
				if file.md5 == "" {
					file.md5, err = Md5(file.path)
					if err != nil {
						return err
					}
				}
				if file.md5 == entryMd5 {
					log.Debug("remove:", path)
					err = os.Remove(path)
					if err != nil {
						return err
					}
					return nil
				}
			}
		}
		return nil
	})
}
