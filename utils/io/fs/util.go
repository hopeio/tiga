package fs

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/hopeio/lemon/utils/crypto"
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
	defer file.Close()
	hash := md5.New()
	_, err = io.Copy(hash, file)
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
	fileName = crypto.EncodeMD5(fileName)
	return fileName + ext
}
