package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/urfave/cli"
	"log"
	"os"
	"strings"
	"test/tpl"
)
var (
	server string
	port string
)

func main() {
	//日志
	logs.SetLogger("console")

	logs.Alert("欢迎使用rpcb工具")

	app := cli.NewApp()
	app.Name = "rpcb"
	app.Version = "1.0.0"
	app.Usage = "Create a rpcx template"

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
				logs.Info("%s 模板创建成功",server)
			}
			//创建文件
			f,err := os.Create(fmt.Sprintf("%s/%s.go",path,server))
			defer f.Close()
			if err != nil {
				logs.Error(err.Error())
			} else {
				_,err = f.Write([]byte(fmt.Sprintf(tpl.Tpl_main,port,server,strFirstToUpper(server))))
				if err != nil {
					logs.Error(err)
				}
			}

			f2,err := os.Create(fmt.Sprintf("%s/handler.go",path))
			defer f2.Close()
			if err != nil {
				logs.Error(err.Error())
			} else {
				_,err = f2.Write([]byte(fmt.Sprintf(tpl.Tpl_handler,strFirstToUpper(server),strFirstToUpper(server))))
				if err != nil {
					logs.Error(err)
				}
			}
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


func strFirstToUpper(str string) string {
	temp := []byte(str)
	return strings.Join([]string{strings.ToUpper(string(temp[0])),string(temp[1:])},"")
}