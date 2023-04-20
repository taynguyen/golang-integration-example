package main

import (
	"context"
	"fmt"
	"taynguyen/sample/src/repo/orm"
	"time"

	"github.com/go-pg/pg/v10"
)

func main() {
	ctx := context.Background()

	db := pg.Connect(&pg.Options{
		Addr:     "localhost:5435",
		User:     "taylor",
		Password: "taylor",
		Database: "taylor",
	})

	if err := db.Ping(ctx); err != nil {
		panic(err)
	}

	b := &orm.Order{
		Id:         1,
		CustomerId: 1,
		Price:      1000,
		CreatedAt:  time.Now(),
	}

	rs, err := db.Model(b).Insert()
	if err != nil {
		panic(err)
	}
	fmt.Println(rs)

}
