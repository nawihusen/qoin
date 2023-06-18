package mysql

import (
	"Qoin/domain"
	"context"
	"errors"

	"database/sql"

	log "github.com/sirupsen/logrus"
)

type mysqlQoinRepository struct {
	Conn *sql.DB
}

func NewMySQLQoinRepository(Conn *sql.DB) domain.QoinMySQLRepository {
	return &mysqlQoinRepository{Conn}
}

func (db *mysqlQoinRepository) InsertCustomer(ctx context.Context, customer domain.Customer) (id int64, err error) {
	query := `INSERT INTO customer (name, table_no) VALUES (?, ?)`
	log.Debug("Query : " + query)

	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, customer.Name, customer.TableNo)
	if err != nil {
		return
	}

	id, err = res.LastInsertId()

	return
}

func (db *mysqlQoinRepository) InsertOrder(ctx context.Context, custumerID int64) (id int64, err error) {
	query := `INSERT INTO orders (customer_id, dtm_crt) VALUES (?, NOW())`
	log.Debug("Query : " + query)

	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, custumerID)
	if err != nil {
		return
	}

	id, err = res.LastInsertId()

	return
}

func (db *mysqlQoinRepository) InsertDetail(ctx context.Context, orderID int64, order *domain.OrderRequst, menu map[int64]domain.Menu) (id int64, err error) {
	query := `INSERT INTO detail (order_id, menu_id, amount, total) VALUES `
	var params []interface{}

	for i, v := range order.MenuID {
		query += ` (?, ?, ?, ?),`
		params = append(params, orderID, v, order.Amount[i], menu[v].Price*order.Amount[i])
	}

	if query[len(query)-1:] == "," {
		query = query[:len(query)-1]
	}

	log.Debug("Query : " + query)

	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, params...)
	if err != nil {
		return
	}

	id, err = res.LastInsertId()

	return
}

func (db *mysqlQoinRepository) SelectMenu(ctx context.Context, menuID int64) (menu domain.Menu, err error) {
	query := `SELECT id, name, stock, price FROM menu WHERE id = ?`
	log.Debug(query)

	row := db.Conn.QueryRowContext(ctx, query, menuID)

	err = row.Scan(&menu.ID, &menu.Name, &menu.Stock, &menu.Price)

	if err != nil {
		if err == sql.ErrNoRows {
			err = errors.New("menu not found")
		}
		return
	}

	return
}

func (db *mysqlQoinRepository) SelectListMenu(ctx context.Context) (menu []domain.Menu, err error) {
	query := `SELECT id, name, stock, price FROM menu`
	log.Debug(query)

	rows, err := db.Conn.QueryContext(ctx, query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var i domain.Menu
		err = rows.Scan(&i.ID, &i.Name, &i.Stock, &i.Price)
		if err != nil {
			return
		}

		menu = append(menu, i)
	}

	return
}

func (db *mysqlQoinRepository) InsertNewMenu(ctx context.Context, menu domain.Menu) (err error) {
	query := `INSERT INTO menu (name, stock, price) VALUES (?, ?, ?)`
	log.Debug("Query : " + query)

	stmt, err := db.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx, menu.Name, menu.Stock, menu.Price)
	if err != nil {
		return
	}

	return
}
func (db *mysqlQoinRepository) SelectOrder(ctx context.Context, from, to string) (orders []domain.Order, err error) {
	query := `SELECT id, customer_id FROM orders WHERE dtm_crt BETWEEN '` + from + `' AND '` + to + `'`
	log.Debug(query)

	rows, err := db.Conn.QueryContext(ctx, query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var i domain.Order
		err = rows.Scan(&i.ID, &i.CustomerID)
		if err != nil {
			return
		}

		orders = append(orders, i)
	}

	return
}

func (db *mysqlQoinRepository) SelectDetails(ctx context.Context, id int64) (detail []domain.Detail, err error) {
	query := `SELECT id, order_id, menu_id, amount, total FROM detail WHERE order_id = ?`
	log.Debug(query)

	rows, err := db.Conn.QueryContext(ctx, query, id)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var i domain.Detail
		err = rows.Scan(&i.ID, &i.OrderID, &i.MenuID, &i.Amount, &i.Total)
		if err != nil {
			return
		}

		detail = append(detail, i)
	}

	return
}
