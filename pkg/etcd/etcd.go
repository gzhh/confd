package etcd

import (
	"context"
	"log"
	"sync"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type Callback func(x, y interface{})

type Etcd struct {
	LoadTimeout time.Duration
	once        sync.Once
	cli         *clientv3.Client
}

func NewEtcd() *Etcd {
	return &Etcd{
		LoadTimeout: time.Second * 5,
	}
}

func (e *Etcd) Init(config clientv3.Config) {
	e.once.Do(func() {
		var err error
		e.cli, err = clientv3.New(config)
		if err != nil {
			panic(err)
		}
	})
}

func (e *Etcd) Load(key string, callback Callback) {
	ctx, cancel := context.WithTimeout(context.Background(), e.LoadTimeout)
	defer cancel()

	resp, err := e.cli.Get(ctx, key)
	if err != nil {
		log.Panic(err)
	}

	for _, ev := range resp.Kvs {
		// log.Printf("Key:%s Value:%s\n", ev.Key, ev.Value)
		if string(ev.Key) == key {
			callback(ev.Key, ev.Value)
		}
	}
}

func (s *Etcd) Watch(key string, callback Callback) {
	go func(cli *clientv3.Client) {
		log.Println("watching...")
		rch := cli.Watch(context.Background(), key)
		for wresp := range rch {
			for _, ev := range wresp.Events {
				// log.Printf("Type:%s Key:%s Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
				if string(ev.Kv.Key) == key {
					callback(ev.Kv.Key, ev.Kv.Value)
				}
			}
		}
	}(s.cli)
}

func (e *Etcd) Close() {
	if err := e.cli.Close(); err != nil {
		log.Panic(err)
	}
}
