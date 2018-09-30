package zk

import (
	"errors"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

type createType struct {
	rpath string
	err   error
}

//Create 创建节点
func (client *ZookeeperClient) Create(path string, data string, flags int32, acl []zk.ACL) (rpath string, err error) {
	if !client.isConnect {
		err = errors.New("链接已关闭")
		return
	}
	if client.done {
		err = errors.New("链接已手动关闭")
		return
	}
	//开启协程创建节点
	ch := make(chan interface{}, 1)
	go func(chan interface{}) {
		data, err := client.conn.Create(path, []byte(data), flags, acl)
		ch <- createType{rpath: data, err: err}
	}(ch)
	select {
	case <-time.After(time.Second * 2):
		err = errors.New("创建节点超时")
		return
	case data := <-ch:
		err = data.(createType).err
		if err != nil {
			return "", err
		}
		rpath = data.(createType).rpath
		return rpath, nil
	}
}
