package repo

import (
	"taynguyen/sample/src/repo/order"

	"github.com/go-pg/pg/v10"
)

type Repo interface {
	Order() order.IRepo
}

func New(db *pg.DB) Repo {
	return &impl{
		order: order.New(db),
	}
}

type impl struct {
	order order.IRepo
}

func (i impl) Order() order.IRepo {
	return i.order
}
