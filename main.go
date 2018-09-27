package main

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"
	zk "github.com/wule61/lib/zookeeper"
)

func main() {
	c, err := zk.New([]string{"localhost:2181"}, time.Second)
	if err != nil {
		logs.Error(err)
	}
	c.Connect()
	defer c.Close()
	value, v, _ := c.Get("/zookeeper/quota")
	fmt.Println(string(value))
	fmt.Println(v)
	//select {}
}
