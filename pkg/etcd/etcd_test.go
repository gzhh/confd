package etcd

import (
	"testing"
	"time"

	"github.com/spf13/cast"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func TestLoad(t *testing.T) {
	etcd := NewEtcd()
	etcd.Init(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: time.Second * 5,
	})

	key := "/config.yaml"
	etcd.Load(key, func(x, y interface{}) {
		t.Log(cast.ToString(y))
	})
}

func TestWatch(t *testing.T) {
	etcd := NewEtcd()
	etcd.Init(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: time.Second * 5,
	})

	key := "/config.yaml"
	etcd.Watch(key, func(x, y interface{}) {
		t.Log(cast.ToString(y))
	})

	time.Sleep(time.Second * 20)
}
