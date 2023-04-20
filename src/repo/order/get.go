package order

import "taynguyen/sample/src/repo/orm"

func (i impl) GetAll() ([]orm.Order, error) {
	var orders []orm.Order
	err := i.db.Model(&orders).Select()
	return orders, err
}
