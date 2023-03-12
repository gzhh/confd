package confd

import (
	"bytes"
	"log"
	"time"

	"github.com/gzhh/confd/pkg/etcd"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type Config struct {
	// Etcd url list
	Endpoint []string

	// The timeout for failing to establish a etcd connection
	DialTimeout time.Duration

	// Etcd configuration file path
	Key string

	// Etcd configuration file content type
	Type string
}

func New(config Config) *viper.Viper {
	viper := viper.New()
	viper.SetConfigType(config.Type)

	etcd := etcd.NewEtcd()
	etcd.Init(clientv3.Config{
		Endpoints:   config.Endpoint,
		DialTimeout: config.DialTimeout,
	})

	etcd.Load(config.Key, func(x, y interface{}) {
		if err := viper.ReadConfig(bytes.NewBufferString(cast.ToString(y))); err != nil {
			log.Panicln(err.Error())
		}
	})

	etcd.Watch(config.Key, func(x, y interface{}) {
		if err := viper.ReadConfig(bytes.NewBufferString(cast.ToString(y))); err != nil {
			log.Panicln(err.Error())
		}
	})

	return viper
}
