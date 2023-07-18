package etcd

import (
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func getEtcdEndpoints() []string {
	return []string{"192.168.1.30:2379"}
}

func GetEtcdClient() (*clientv3.Client, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   getEtcdEndpoints(),
		DialTimeout: 2 * time.Second,
	})
	return cli, err
}
