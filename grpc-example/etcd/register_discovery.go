package etcd

import (
	"context"
	"errors"
	"fmt"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func CusServiceRegister(serviceName, addr string) error {
	cli, _ := GetEtcdClient()

	ctx := context.Background()

	// 创建租约
	leaseRes, err := cli.Grant(ctx, 10)
	if err != nil {
		return err
	}

	// 向 etcd 写数据
	_, err = cli.Put(ctx, serviceName, addr, clientv3.WithLease(leaseRes.ID))
	if err != nil {
		return err
	}

	keepaliveCh, err := cli.KeepAlive(ctx, leaseRes.ID)
	if err != nil {
		return err
	}

	go func() {
		for item := range keepaliveCh {
			fmt.Printf("leaseID: %x, ttl: %v\n", item.ID, item.TTL)
		}
	}()

	return nil
}

func CusServiceDiscovery(serviceName string) (string, error) {
	cli, _ := GetEtcdClient()

	ctx := context.Background()

	res, _ := cli.Get(ctx, serviceName)

	for _, item := range res.Kvs {
		// fmt.Printf("item: %v", item)
		if string(item.Key) == serviceName {
			return string(item.Value), nil
		}
	}

	return "", errors.New(fmt.Sprintf("service %s not found.", serviceName))
}
