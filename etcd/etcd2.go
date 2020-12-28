package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

func main() {
	cli, _ := clientv3.New(clientv3.Config{
		Endpoints: []string{"localhost:2379"},
		// Endpoints: []string{"localhost:2379", "localhost:22379", "localhost:32379"}
		DialTimeout: 5 * time.Second,
	})
	kv := clientv3.NewKV(cli)
	putResp, err := kv.Put(context.TODO(), "/test/key1", "Hello etcd!")
	//putResp, err := cli.Put(context.Background(), "test1", "123")
	getResp, err := kv.Get(context.TODO(), "/test/key1")
	fmt.Println(putResp, getResp, err)
}
