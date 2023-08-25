package mysql

import (
	timei "github.com/hopeio/lemon/utils/time"
	"time"
)

func Now() string {
	return time.Now().Format(timei.TimeFormat)
}
