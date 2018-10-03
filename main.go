package main

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"
	zk "github.com/wule61/lib/zookeeper"
)

func main() {
	c, err := zk.New([]string{"127.0.0.1:2181", "127.0.0.1:2182", "127.0.0.1:2183"}, time.Second)
	if err != nil {
		logs.Error(err)
	}
	c.Connect()
	defer c.Close()
	snapshots, errs := c.WatchServerList("/node")
	go func() {
		for {
			select {
			case serverList := <-snapshots:
				fmt.Println(serverList)
			case erros := <-errs:
				fmt.Println(erros)
			}
		}
	}()

	configs, errors := c.WatchGetData("/node")

	go func() {
		for {
			select {
			case configData := <-configs:
				fmt.Println(string(configData))
			case err = <-errors:
				fmt.Println(err)
			}
		}
	}()
	select {}
}
