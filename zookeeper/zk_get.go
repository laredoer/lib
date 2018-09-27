package zk

import "errors"

//Get 获取节点的值
func (client *ZookeeperClient) Get(path string) (value []byte, version int32, err error) {
	if !client.isConnect {
		err = errors.New("链接已断开，获取节点失败")
		return
	}
	if client.done {
		err = errors.New("链接已经手动关闭，获取节点失败")
	}
	value, stat, err := client.conn.Get(path)
	version = stat.Version
	return
}
