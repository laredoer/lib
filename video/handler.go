package main

import "context"

type Video struct {
	db *Database
}

func (v *Video) GetRepo() Repository {
	return &VideoRepository{rp: v.db}
}

func (v *Video) Get(ctx context.Context, args Request, reply *Response) (err error) {
	err = v.GetRepo().Create()
	if err != nil {
		return err
	}
	return nil
}
