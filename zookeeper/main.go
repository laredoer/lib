package zk

import (
	"sync"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

type ZookeeperClient struct {
	servers   []string        //zk地址
	timeout   time.Duration   //超时时间
	conn      *zk.Conn        //zk连接实例
	eventChan <-chan zk.Event //zk事件
	useCount  int
	isConnect bool //是否连接
	once      sync.Once
	CloseCh   chan struct{} //关闭通道
	done      bool          //是否手动关闭
}

//创建Zookeeper实例
func New(servers []string, timeout time.Duration) (*ZookeeperClient, error) {
	client := &ZookeeperClient{servers: servers, timeout: timeout, useCount: 0}
	client.CloseCh = make(chan struct{})
	return client, nil
}
