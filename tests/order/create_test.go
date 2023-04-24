package order

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"taynguyen/sample/src/cmd/router"
	"taynguyen/sample/src/handler"
	"taynguyen/sample/src/repo"
	"taynguyen/sample/src/repo/orm"
	"taynguyen/sample/src/utils"
	"testing"

	"github.com/go-pg/pg/v10"
	"github.com/stretchr/testify/suite"
)

type OrderSuite struct {
	suite.Suite

	db *pg.DB

	srv *httptest.Server
}

func (s *OrderSuite) SetupSuite() {
	// println("SetupSuite")
	// Start server
	ctx := context.Background()
	if err := utils.InitializeSnowFlake(); err != nil {
		panic(err)
	}

	s.db = pg.Connect(&pg.Options{
		Addr:     "localhost:5435",
		User:     "taylor",
		Password: "taylor",
		Database: "taylor",
	})
	if err := s.db.Ping(ctx); err != nil {
		panic(err)
	}

	repo := repo.New(s.db)
	h := handler.New(repo)
	s.srv = httptest.NewServer(router.CreateRouter(h))
}

func (s *OrderSuite) TearDownSuite() {
	// Shutdown server
	if s.srv != nil {
		s.srv.Close()
	}
}

func (s *OrderSuite) SetupTest() {
	// println("SetupTest")
	// Clean data
	s.db.Exec(`DELETE FROM orders`)
}

func (s *OrderSuite) TearDownTest() {
	// println("TearDownTest")
}

func (s *OrderSuite) TestPing() {
	res, err := http.Get(fmt.Sprintf("%s/ping", s.srv.URL))
	if err != nil {
		s.Error(err, "Error when call ping")
	}
	body, err := io.ReadAll(res.Body)
	s.Assertions.Nil(err, "Error when read body")
	s.Assertions.Equal(`{"message":"pong"}`, string(body), "Body should be pong")
}

func (s *OrderSuite) Test_Create_Valid_Order() {
	// WHEN create order
	order1 := &handler.NewOrderRq{
		Title:      "Order 1",
		CustomerID: 1,
		Price:      100,
	}
	buf, _ := json.Marshal(order1)
	rqIo := bytes.NewReader(buf)
	res, err := http.Post(fmt.Sprintf("%s/order", s.srv.URL), "application/json", rqIo)

	// THEN validate response
	s.Equal(http.StatusOK, res.StatusCode, "Status code should be 200")
	s.Nil(err, "Error when call create order")

	body, err := io.ReadAll(res.Body)
	s.Nil(err, "Error when read body")
	o := &handler.OrderResponse{}
	s.Nil(json.Unmarshal(body, o), "Error when unmarshal body")
	s.Equal(order1.CustomerID, o.CustomerID, "CustomerID")
	s.Equal(order1.Title, o.Title, "Title")
	s.Equal(order1.Price, o.Price, "Price")
	s.True(o.ID > 0, "ID should be greater than 0")
	s.True(o.CreatedAt != "", "CreatedAt should not be empty")
}

func (s *OrderSuite) Test_Create_Invalid_JSON_Should_Return_4xx() {
	// WHEN create order
	rqIo := bytes.NewReader([]byte(`{ invalid json here`))
	res, err := http.Post(fmt.Sprintf("%s/order", s.srv.URL), "application/json", rqIo)

	// THEN validate response
	s.Equal(http.StatusBadRequest, res.StatusCode, "Status code should be 400")
	s.Nil(err, "Error when call create order")
}

func (s *OrderSuite) Test_GetAll_Should_Return_All_Order() {
	// GIVEN 2 orders in DB
	order1 := &orm.Order{
		Id:         1,
		Title:      "Order 1",
		CustomerId: 1,
		Price:      100,
	}
	_, err := s.db.Model(order1).Insert()
	s.Nil(err, "Error when insert order 1")

	order2 := &orm.Order{
		Id:         2,
		Title:      "Order 2",
		CustomerId: 2,
		Price:      200,
	}
	_, err = s.db.Model(order2).Insert()
	s.Nil(err, "Error when insert order 2")

	// WHEN get all orders
	res, err := http.Get(fmt.Sprintf("%s/orders", s.srv.URL))

	// THEN validate response
	s.Equal(http.StatusOK, res.StatusCode, "Status code should be 200")
	s.Nil(err, "Error when call create order")

	body, err := io.ReadAll(res.Body)
	s.Nil(err, "Error when read body")
	orders := []*handler.OrderResponse{}
	s.Nil(json.Unmarshal(body, &orders), "Error when unmarshal body")
	s.Equal(2, len(orders), "Should return 2 orders")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestOrderSuite(t *testing.T) {
	suite.Run(t, new(OrderSuite))
}
