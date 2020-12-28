package main

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/net/context"

	"github.com/coreos/etcd/clientv3"
)

var (
	dialTimeout    = 5 * time.Second
	requestTimeout = 2 * time.Second
	endpoints      = []string{"localhost:2379"}
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: dialTimeout,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	go func() {
		rch := cli.Watch(context.Background(), "", clientv3.WithPrefix())
		for wresp := range rch {
			for _, ev := range wresp.Events {
				fmt.Printf("Watch: %s %q: %q \n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}
		}
	}()

	kvs := map[string]string{
		"key1":  "value1",
		"key10": "value10",
		"key5":  "value5",
		"keyk":  "valuek",
		"key2":  "value2",
	}
	for k, v := range kvs {
		cli.Put(context.TODO(), k, v)
	}

	if resp, err := cli.Get(context.TODO(), "key1", clientv3.WithRange("keyk")); err != nil {
		log.Fatal(err)
	} else {
		log.Println(resp)
	}
}
