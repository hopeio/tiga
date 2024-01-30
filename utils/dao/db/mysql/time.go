package mysql

import (
	timei "github.com/hopeio/tiga/utils/time"
	"time"
)

func Now() string {
	return time.Now().Format(timei.TimeFormat)
}
