package order

import (
	"taynguyen/sample/src/repo/orm"

	"github.com/go-pg/pg/v10"
)

type IRepo interface {
	Create(o orm.Order) (orm.Order, error)

	GetAll() ([]orm.Order, error)
}

type impl struct {
	db pg.DBI
}

func New(db *pg.DB) IRepo {
	return &impl{db: db}
}
