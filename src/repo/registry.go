package repo

import (
	"taynguyen/sample/src/repo/order"

	"github.com/go-pg/pg/v10"
)

type Repo interface {
	Order() order.IRepo

	// DoInTx(fn func(repo *pg.Tx) error) error
}

func New(db *pg.DB) Repo {
	return &impl{
		db:    db,
		order: order.New(db),
	}
}

type impl struct {
	db *pg.DB

	order order.IRepo
}

func (i impl) Order() order.IRepo {
	return i.order
}
