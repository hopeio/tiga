package mysql

import (
	"fmt"
	pkdb "github.com/hopeio/tiga/initialize/basic_dao/gormdb"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseConfig pkdb.DatabaseConfig

func (conf *DatabaseConfig) Init() {
	(*pkdb.DatabaseConfig)(conf).Init()
}

func (conf *DatabaseConfig) Build() *gorm.DB {
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		conf.User, conf.Password, conf.Host,
		conf.Port, conf.Database, conf.Charset)
	return (*pkdb.DatabaseConfig)(conf).Build(mysql.Open(url))
}

type DB pkdb.DB

func (db *DB) Config() any {
	return (*DatabaseConfig)(&db.Conf)
}

func (db *DB) SetEntity(entity interface{}) {
	db.DB = (*DatabaseConfig)(&db.Conf).Build()
}

func (db *DB) Close() error {
	return nil
}
