package main

import (
	"fmt"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

func main() {
	conn, err := getConnect()
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	conn.Create("/go_servers", nil, 0, zk.WorldACL(zk.PermAll))

	time.Sleep(20 * time.Second)
}

func getConnect() (conn *zk.Conn, err error) {
	var hosts = []string{"localhost:2181"} //server端host
	conn, _, err = zk.Connect(hosts, time.Second*5)
	if err != nil {
		fmt.Println("错误:", err)
		return nil, err
	}
	return conn, err
}
