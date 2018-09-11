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
	// init db
	DB.Init()
	defer DB.Close()
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

func (v *%s) Get(ctx context.Context,args Request,reply *Response) (err error){
	return nil
}
`
const Tpl_datastore = `
package main

import (
	"flag"
	"fmt"

	"github.com/astaxie/beego/logs"
	// MySQL driver.
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	username = flag.String("username", "root", "db username")
	password = flag.String("password", "root", "password")
	dbAddr   = flag.String("abAddr", "localhost:3306", "db addr")
	name     = flag.String("name", "video", "db name")
)

type Database struct {
	Self *gorm.DB
}

var DB *Database

func openDB(username, password, addr, name string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		//"Asia/Shanghai"),
		"Local")

	db, err := gorm.Open("mysql", config)
	if err != nil {
		logs.Error(err, "Database connection failed. Database name: %s", name)
	}

	// set for db connection
	setupDB(db)

	return db
}

func setupDB(db *gorm.DB) {
	//db.DB().SetMaxOpenConns(20000) // 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	db.DB().SetMaxIdleConns(0) // 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
}

// used for cli
func InitSelfDB() *gorm.DB {
	return openDB(*username, *password, *dbAddr, *name)
}

func GetSelfDB() *gorm.DB {
	return InitSelfDB()
}

func (db *Database) Init() {
	DB = &Database{
		Self: GetSelfDB(),
	}
}

func (db *Database) Close() {
	DB.Self.Close()
}
`
const Tpl_repository = `
package main

type Repository interface {
	Create() error
}

type %sModel struct {
	Username string
	Password string
}

// 接口实现
func (u *%sModel) Create() error {
	//TODO
	return nil
}`
