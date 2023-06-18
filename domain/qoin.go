package domain

import (
	"context"
	"time"
)

type Customer struct {
	ID      int64
	Name    string
	TableNo int
}

type Order struct {
	ID         int64
	CustomerID int64
	DtmCrt     time.Time
}

type Menu struct {
	ID    int64
	Name  string `json:"name" form:"name" validate:"required"`
	Stock int    `json:"stock" form:"stock" validate:"required"`
	Price int    `json:"price" form:"price" validate:"required"`
}

type Detail struct {
	ID      int64
	OrderID int64
	MenuID  int64
	Amount  int
	Total   int
}

type OrderRequst struct {
	Name    string  `json:"name" form:"name" validate:"required"`
	TableNo int     `json:"table_no" form:"table_no" validate:"required"`
	MenuID  []int64 `json:"menu_id" form:"menu_id" validate:"required"`
	Amount  []int   `json:"amount" form:"amount" validate:"required"`
}

type MenuResponse struct {
	DetailID int64     `json:"detail_id" form:"detail_id"`
	TotalAll int       `json:"total_all" form:"total_all"`
	Receipt  []Receipt `json:"receipt" form:"receipt"`
}

type Receipt struct {
	Name  string `json:"name" form:"name"`
	Price int    `json:"price" form:"price"`
	Total int    `json:"total" form:"total"`
}

type Stock struct {
	Stock int `json:"stock" form:"stock"`
}

type IncomeRequest struct {
	From time.Time     `json:"from" form:"from"`
	To   time.Duration `json:"to" form:"to"`
}

type QoinUsecase interface {
	MakeOrder(ctx context.Context, order *OrderRequst) (response MenuResponse, err error)
	GetReceipt(ctx context.Context, detailID int64) (response MenuResponse, err error)
	GetIncome(ctx context.Context, from, to string) (income int64, err error)
	GetStock(ctx context.Context, menuID int64) (stock Stock, err error)
	GetListMenu(ctx context.Context) (menu []Menu, err error)
	AddNewMenu(ctx context.Context, menu Menu) (err error)
}

type QoinMySQLRepository interface {
	InsertCustomer(ctx context.Context, customer Customer) (id int64, err error)
	InsertOrder(ctx context.Context, custumerID int64) (id int64, err error)
	InsertDetail(ctx context.Context, orderID int64, orderReq *OrderRequst, menu map[int64]Menu) (id int64, err error)
	SelectListMenu(ctx context.Context) (menu []Menu, err error)
	SelectMenu(ctx context.Context, menuID int64) (menu Menu, err error)
	InsertNewMenu(ctx context.Context, menu Menu) (err error)
	SelectOrder(ctx context.Context, from, to string) (orders []Order, err error)
	SelectDetails(ctx context.Context, orderID int64) (detail []Detail, err error)
}
