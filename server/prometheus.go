package server

import (
	"github.com/hashicorp/go-metrics/prometheus"
	"github.com/hopeio/lemon/initialize"
)

var (
	metric *metrics.Metrics
)

func init() {
	sink, _ := prometheus.NewPrometheusSink()
	conf := metrics.DefaultConfig(initialize.GlobalConfig.Module)
	metric, _ = metrics.New(conf, sink)
	metric.EnableHostnameLabel = true
}
