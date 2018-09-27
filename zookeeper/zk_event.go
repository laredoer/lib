package zk

import (
	"github.com/astaxie/beego/logs"
	"github.com/samuel/go-zookeeper/zk"
)

func (client *ZookeeperClient) eventWatch() {
START:
	for {
		select {
		case <-client.CloseCh:
			break START
		case v, ok := <-client.eventChan:
			if ok {
				switch v.State {
				case zk.StateAuthFailed:
					logs.Info("链接失败")
					client.isConnect = false
				// 已经连接成功
				case zk.StateConnected:
					logs.Alert("链接成功")
					client.isConnect = true
				// 连接Session失效
				case zk.StateExpired:
					logs.Info("连接Session失效")
					client.isConnect = false
				// 网络连接不成功
				case zk.StateDisconnected:
					logs.Warn("zk已断开连接:%v", client.servers)
					client.isConnect = false
				// 网络断开，正在连接
				case zk.StateConnecting:
					logs.Warn("网络断开，正在连接")
					client.isConnect = false
				case zk.StateHasSession:
					client.isConnect = true
				}
			} else {
				logs.Warn("网络断开")
				client.isConnect = false
				break START
			}
		}
	}
}
