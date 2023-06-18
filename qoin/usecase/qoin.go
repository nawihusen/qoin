package usecase

import (
	"Qoin/domain"
	"context"
	"fmt"
)

type qoinUsecase struct {
	qoinMySQLRepo domain.QoinMySQLRepository
}

// NewAccountUsecase is constructor of account usecase
func NewQoinUsecase(qoinMySQLRepo domain.QoinMySQLRepository) domain.QoinUsecase {
	return &qoinUsecase{
		qoinMySQLRepo: qoinMySQLRepo,
	}
}

func (u *qoinUsecase) MakeOrder(ctx context.Context, order *domain.OrderRequst) (response domain.MenuResponse, err error) {
	customerID, err := u.qoinMySQLRepo.InsertCustomer(ctx, domain.Customer{Name: order.Name, TableNo: order.TableNo})
	if err != nil {
		return
	}

	orderID, err := u.qoinMySQLRepo.InsertOrder(ctx, customerID)
	if err != nil {
		return
	}

	list, err := u.qoinMySQLRepo.SelectListMenu(ctx)

	menu := map[int64]domain.Menu{}
	for _, v := range list {
		menu[v.ID] = domain.Menu{Name: v.Name, Price: v.Price}
	}

	detailID, err := u.qoinMySQLRepo.InsertDetail(ctx, orderID, order, menu)
	if err != nil {
		return
	}

	response.DetailID = detailID
	response.TotalAll = 0

	for i, v := range order.MenuID {
		var temp domain.Receipt
		temp.Name = menu[v].Name
		temp.Price = menu[v].Price
		temp.Total = menu[v].Price * order.Amount[i]

		response.TotalAll += temp.Total

		response.Receipt = append(response.Receipt, temp)
	}

	return
}

func (u *qoinUsecase) GetStock(ctx context.Context, menuID int64) (stock domain.Stock, err error) {
	menu, err := u.qoinMySQLRepo.SelectMenu(ctx, menuID)
	if err != nil {
		return
	}

	stock.Stock = menu.Stock

	return
}

func (u *qoinUsecase) GetListMenu(ctx context.Context) (menu []domain.Menu, err error) {
	menu, err = u.qoinMySQLRepo.SelectListMenu(ctx)

	return
}

func (u *qoinUsecase) AddNewMenu(ctx context.Context, menu domain.Menu) (err error) {
	err = u.qoinMySQLRepo.InsertNewMenu(ctx, menu)

	return
}

func (u *qoinUsecase) GetReceipt(ctx context.Context, detailID int64) (response domain.MenuResponse, err error) {
	return
}

func (u *qoinUsecase) GetIncome(ctx context.Context, from, to string) (income int64, err error) {
	fmt.Println(from)
	fmt.Println(to)
	order, err := u.qoinMySQLRepo.SelectOrder(ctx, from, to)
	if err != nil {
		return
	}
	fmt.Println(order)

	for _, v := range order {
		detail, er := u.qoinMySQLRepo.SelectDetails(ctx, v.ID)
		if er != nil {
			return income, er
		}

		for _, v2 := range detail {
			income += int64(v2.Total)
		}

	}

	return
}
