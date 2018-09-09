package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/urfave/cli"
	"log"
	"os"
	"test/tpl"
)

func main() {
	//日志
	logs.SetLogger("console")

	logs.Error("欢迎使用rpcb工具")

	app := cli.NewApp()
	app.Name = "rpcb"
	app.Version = "1.0.0"
	app.Usage = "Create a rpcx template"
	var server string
	var port string
	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "server",
			Value: "example",
			Usage: "create a server template",
			Destination: &server,         //取到的FLAG值，赋值到这个变量
		},
		cli.StringFlag{
			Name: "port",
			Value:"8973",
			Usage:"server port",
			Destination: &port,
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.String("server") != "" {

			path := fmt.Sprintf("./%s",server)
			b , err := PathExists(path)
			if err != nil || b == true {
				logs.Error("默认创建example,文件: %s 已经存在", path)
				return err
			}
			// 创建文件夹
			err = os.Mkdir(path, os.ModePerm)
			if err != nil {
				logs.Error("创建失败:%s",err)
			} else {
				logs.Info("创建成功")
			}

			f,err := os.Create(fmt.Sprintf("%s/%s.go",path,server))
			defer f.Close()
			if err != nil {
				logs.Error(err.Error())
			} else {
				_,err = f.Write([]byte(fmt.Sprintf(tpl.Tpl_main,port,server)))
				if err != nil {
					logs.Error(err)
				}
			}
		} else {
			fmt.Println("Hello", server)
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}