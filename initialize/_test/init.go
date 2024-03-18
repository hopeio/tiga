package main

import (
	"fmt"
	"github.com/hopeio/tiga/initialize"
)

type TestConfig struct {
	A GroupA
	B GroupB
}

func (t *TestConfig) Init() {

}

type GroupA struct {
	AString string
	AInt    int
}

type GroupB struct {
	BString string
	BInt    int
}

func main() {
	conf := &TestConfig{}
	initialize.Start(conf, nil)
	fmt.Println(initialize.GlobalConfig.BasicConfig)
	fmt.Println(conf.B)
}
