package server

import "github.com/hopeio/tiga/server"

type ServerConfig server.ServerConfig

func (c *ServerConfig) Init() {
	(*server.ServerConfig)(c).Init()
}

func (c *ServerConfig) Origin() *server.ServerConfig {
	return (*server.ServerConfig)(c)
}
