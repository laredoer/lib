package main

// 定义请求结构体,根据实际情况做修改
type Request struct {
	UserId   int64
	UserName string
	PassWord string
}

//Response 结构体,根据实际情况做修改
type Response struct {
	Errcode int
	Errmsg  string
	Data    interface{}
}
