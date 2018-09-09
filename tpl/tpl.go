package tpl


const Tpl_main = `package main

import (
	"flag"
	"fmt"
	"github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
	"log"
	"time"
)

var (
	addr  = flag.String("addr","localhost:%s","server addr")
	consulAddr = flag.String("consulAddr","132.232.109.253:8500","consul addr")
	basePath = flag.String("basePath","/qu_video","prefix path")
)
func main() {
	flag.Parse()
	s := server.NewServer()
	addRegistryPlugin(s)
	s.RegisterName("%s",new(%s),"")
	err := s.Serve("tcp",*addr)
	if err != nil {
		fmt.Println(err)
	}
}
// 服务注册
func addRegistryPlugin(s *server.Server) {
	r := &serverplugin.ConsulRegisterPlugin{
		ServiceAddress: "tcp@" + *addr,
		ConsulServers: []string{*consulAddr},
		BasePath: *basePath,
		Metrics: metrics.NewRegistry(),
		UpdateInterval: time.Second * 10,
	}
	err := r.Start()
	if err != nil {
		log.Fatal(err)
	}
	s.Plugins.Add(r)
}`

const Tpl_handler = `package main

import "context"

type %s struct {
}

// 定义请求结构体,根据实际情况做修改
type Request struct {
	UserId int64 	
	UserName string 
	PassWord string 
}

//Response 结构体,根据实际情况做修改
type Response struct {
	Errcode int         
	Errmsg  string     
	Data    interface{} 
}

func (v *%s) Get(ctx *context.Context,args Request,reply *Response) (err error){
	return nil
}
`
