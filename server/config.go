package server

import (
	gini "github.com/hopeio/lemon/utils/net/http/gin"
	"github.com/rs/cors"
	"net/http"
	"time"
)

type ServerConfig struct {
	Protocol                        string
	Domain                          string
	Port                            int32
	StaticFs                        []*StaticFsConfig
	ReadTimeout                     time.Duration `expr:"$+5"`
	WriteTimeout                    time.Duration `expr:"$+5"`
	OpenTracing, Prometheus, GenDoc bool
	Gin                             *gini.Config
	GrpcWeb                         bool
	Http3                           *Http3Config
	Cors                            cors.Options
}

func (c *ServerConfig) Init() {
	if c.Port == 0 {
		c.Port = 8080
	}
	c.ReadTimeout = c.ReadTimeout * time.Second
	c.WriteTimeout = c.WriteTimeout * time.Second

	if len(c.Cors.AllowedMethods) == 0 {
		c.Cors.AllowedMethods = []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		}
	}

	if len(c.Cors.AllowedHeaders) == 0 {
		c.Cors.AllowedHeaders = []string{"*"}
	}
}

func defaultServerConfig() *ServerConfig {
	return &ServerConfig{
		Port: 8080,
	}
}

type Http3Config struct {
	Address  string
	CertFile string
	KeyFile  string
}

type StaticFsConfig struct {
	Prefix string
	Root   string
}
