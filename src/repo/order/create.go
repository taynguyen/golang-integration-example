package order

import (
	"taynguyen/sample/src/repo/orm"
	"taynguyen/sample/src/utils"
)

func (i impl) Create(o orm.Order) (orm.Order, error) {
	// i.db.RunInTransaction(context.Background(), func(tx *pg.Tx) error {
	// 	return nil
	// })
	o.Id = utils.GenerateSnowflakeID()
	_, err := i.db.Model(&o).Insert()
	return o, err
}
