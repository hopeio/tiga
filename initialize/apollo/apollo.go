package apollo

import (
	"github.com/hopeio/lemon/utils/configor/apollo"
	"github.com/hopeio/lemon/utils/log"
)

type ApolloConfig apollo.Config

func (conf *ApolloConfig) Build() *apollo.Client {
	//初始化更新配置，这里不需要，开启实时更新时初始化会更新一次
	client, err := (*apollo.Config)(conf).NewClient()
	if err != nil {
		log.Fatal(err)
	}
	return client
}

type ApolloClient struct {
	*apollo.Client
	Conf ApolloConfig
}

func (a *ApolloClient) Config() any {
	return &a.Conf
}

func (a *ApolloClient) SetEntity() {
	a.Client = a.Conf.Build()
}

func (c *ApolloClient) Close() error {
	return c.Client.Close()
}
