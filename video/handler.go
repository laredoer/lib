package main

import "context"

type Video struct {
	db *Database
}

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

func (h *Video) GetRepo() Repository {
	return &VideoRepository{rp: h.db}
}

func (v *Video) Get(ctx context.Context, args Request, reply *Response) (err error) {
	return nil
}
