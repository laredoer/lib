package main

type Repository interface {
	Create() (err error)
}

type VideoRepository struct {
	rp *Database
}

func NewVideoRepository(db *Database) *VideoRepository {
	return &VideoRepository{
		rp: db,
	}
}

// 接口实现
func (v *VideoRepository) Create() (err error) {
	return nil
}
