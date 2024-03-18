package gormi

import (
	contexti "github.com/hopeio/tiga/utils/context"
	"gorm.io/gorm"
)

func NewTraceDB(db *gorm.DB, traceId string) *gorm.DB {
	return db.Session(&gorm.Session{Context: contexti.SetTranceId(traceId), NewDB: true})
}
