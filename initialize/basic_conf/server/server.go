package server

import "github.com/hopeio/lemon/server"

type ServerConfig server.ServerConfig

func (c *ServerConfig) Init() {
	(*server.ServerConfig)(c).Init()
}

func (c *ServerConfig) Origin() *server.ServerConfig {
	return (*server.ServerConfig)(c)
}
