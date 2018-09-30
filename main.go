package main

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"
	zk "github.com/wule61/lib/zookeeper"
)

func main() {
	c, err := zk.New([]string{"132.232.109.253:2181"}, time.Second)
	if err != nil {
		logs.Error(err)
	}
	c.Connect()
	defer c.Close()
	err = c.Delete("/a")
	fmt.Println(err)

	//select {}
}
