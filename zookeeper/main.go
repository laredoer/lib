package main

import (
	"fmt"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

var hosts = []string{"localhost:2181"}
var path = "/watch"
var flags int32 = zk.FlagEphemeral
var data1 = []byte("hello,this is a zk go")
var acls = zk.WorldACL(zk.PermAll)

func main() {
	option := zk.WithEventCallback(callback)
	conn, _, err := zk.Connect(hosts, time.Second*5, option)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	create(conn, path, data1)

	_, _, ech, err := conn.ExistsW(path)
	if err != nil {
		fmt.Println(err)
		return
	}
START:
	for {
		select {
		case _, ok := <-ech:
			if ok {
				break START
			} else {
				fmt.Println("节点不存在")
			}

		}
	}

}

func watchCreataNode(ech <-chan zk.Event) {
	event := <-ech
	fmt.Println("********channel***********")
	fmt.Println("path:", event.Path)
	fmt.Println("type:", event.Type.String())
	fmt.Println("state:", event.State.String())
	fmt.Println("-------------------")
}

func callback(event zk.Event) {
	fmt.Println("********callback***********")
	fmt.Println("path:", event.Path)
	fmt.Println("type:", event.Type.String())
	fmt.Println("state:", event.State.String())
	fmt.Println("-------------------")
}

//创建节点
func create(conn *zk.Conn, path string, data []byte) {
	_, err_create := conn.Create(path, data, flags, acls)
	if err_create != nil {
		fmt.Println(err_create)
		return
	}
}
