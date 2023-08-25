package engine_old

import (
	"time"
)

type Worker[KEY comparable, T, W any] struct {
	Id     uint
	Kind   Kind
	taskCh chan *BaseTask[KEY, T]
	Props  W
}

// WorkStatistics worker统计数据
type WorkStatistics struct {
	averageTimeCost                                time.Duration
	taskDoneCount, taskTotalCount, taskFailedCount uint64
}

// BaseEngineStatistics 基本引擎统计数据
type BaseEngineStatistics = WorkStatistics
