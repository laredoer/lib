package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
)

var (
	addr       = flag.String("addr", "localhost:8972", "server addr")
	consulAddr = flag.String("consulAddr", "132.232.109.253:8500", "consul addr")
	basePath   = flag.String("basePath", "/qu_video", "prefix path")
)

func main() {
	flag.Parse()
	// init db
	DB.Init()
	defer DB.Close()
	s := server.NewServer()
	addRegistryPlugin(s)
	s.RegisterName("video", &Video{db: DB}, "")
	err := s.Serve("tcp", *addr)
	if err != nil {
		fmt.Println(err)
	}
}

// 服务注册
func addRegistryPlugin(s *server.Server) {
	r := &serverplugin.ConsulRegisterPlugin{
		ServiceAddress: "tcp@" + *addr,
		ConsulServers:  []string{*consulAddr},
		BasePath:       *basePath,
		Metrics:        metrics.NewRegistry(),
		UpdateInterval: time.Second * 10,
	}
	err := r.Start()
	if err != nil {
		log.Fatal(err)
	}
	s.Plugins.Add(r)
}
