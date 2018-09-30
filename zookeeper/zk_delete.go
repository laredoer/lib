package zk

import (
	"errors"
	"time"
)

//Delete 删除节点
func (client *ZookeeperClient) Delete(path string) (err error) {
	if !client.isConnect {
		err = errors.New("链接已经关闭，判断节点是否存在失败")
	}
	if client.done {
		err = errors.New("链接已经手动关闭")
	}

	//开启协程删除节点
	ch := make(chan error, 1)
	go func(chan error) {
		ch <- client.conn.Delete(path, -1) //有子节点是不能删除的
	}(ch)
	select {
	case <-time.After(time.Second * 2):
		err = errors.New("删除节点超时")
		return
	case err = <-ch:
		return
	}
}
