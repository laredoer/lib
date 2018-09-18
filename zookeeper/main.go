package main

import (
	"fmt"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

func main() {
	var hosts = []string{"localhost:2181"} //server端host
	conn, _, err := zk.Connect(hosts, time.Second*5)
	if err != nil {
		fmt.Println("错误:", err)
		return
	}
	defer conn.Close()
}
