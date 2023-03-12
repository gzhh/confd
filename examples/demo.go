package main

import (
	"fmt"
	"time"

	"github.com/gzhh/confd"
)

func main() {
	confd := confd.New(confd.Config{
		Endpoint:    []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: time.Second * 5,
		Key:         "/config.yaml",
		Type:        "yaml",
	})

	for i := 0; i < 5; i++ {
		fmt.Println(confd.Get("mysql"))
		time.Sleep(time.Second * 3)
	}
}
