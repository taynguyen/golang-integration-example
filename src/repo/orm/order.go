package orm

import (
	"time"

	"github.com/go-pg/pg/v10"
)

type Order struct {
	Id         int64
	Title      string
	CustomerId int
	Price      int
	CreatedAt  time.Time
	UpdatedAt  pg.NullTime
}

// func (b Booking) String() string {
// 	return fmt.Sprintf("Booking<%d %d %d %s %s>", b.Id, b.CustomerId, b.Price, b.CreatedAt.String(), b.UpdatedAt.String())
// }
