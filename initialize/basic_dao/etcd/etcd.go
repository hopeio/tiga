package etcd

import (
	"context"
	"github.com/hopeio/tiga/initialize"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type EtcdConfig clientv3.Config

func (conf *EtcdConfig) Build() *clientv3.Client {
	client, _ := clientv3.New((clientv3.Config)(*conf))
	resp, _ := client.Get(context.Background(), initialize.InitKey)
	initialize.GlobalConfig.UnmarshalAndSet(resp.Kvs[0].Value)
	return client
}

type Eecd struct {
	*clientv3.Client
	Conf EtcdConfig
}

func (e *Eecd) Config() any {
	return &e.Conf
}

func (e *Eecd) SetEntity() {
	e.Client = e.Conf.Build()
}

func (e *Eecd) Close() error {
	if e.Client == nil {
		return nil
	}
	return e.Client.Close()
}
