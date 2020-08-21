package mysql

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"lbbs-service/domain"
)

func (r mysqlCartRepository) CountOrdersInCart(cartID int) (int, error) {
	query := `
		SELECT count(*) AS count
		FROM orders
		WHERE cart_id = ?;`

	var count int
	err := r.db.QueryRow(query, cartID).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r mysqlCartRepository) FindCartByID(cartID int) (domain.Cart, error) {
	var cartTable cartTable
	cartQuery := "SELECT id, pos_terminal_id FROM carts WHERE id = ?;"
	if err := r.db.Get(&cartTable, cartQuery, cartID); err != nil {
		if err != sql.ErrNoRows {
			log.Printf("couldn't query from cart: %v", err)
		}
		return domain.Cart{}, err
	}
	var cart = domain.Cart{
		ID:            cartTable.ID,
		PosTerminalID: cartTable.PosTerminalID,
	}
	orders, err := r.GetOrders(cart.ID)
	if err == sql.ErrNoRows || err == nil {
		cart.Orders = orders
		return cart, nil
	}
	return domain.Cart{}, err
}

func (r mysqlCartRepository) CreateNewCart(posID int) (domain.Cart, error) {
	log.Error(posID)
	query := "INSERT INTO carts (pos_terminal_id) VALUE (?);"
	tx := r.db.MustBegin()
	if _, err := tx.Exec(query, posID); err != nil {
		log.Printf("couldn't insert to cart %v", err)
		tx.Rollback()
		return domain.Cart{}, err
	}
	tx.Commit()
	return r.FindCartByPosID(posID)
}

func (r mysqlCartRepository) FlushCart(cart domain.Cart) error {
	tx := r.db.MustBegin()
	defer tx.Rollback()
	var orderIDs []int
	for _, o := range cart.Orders {
		orderIDs = append(orderIDs, o.ID)
	}
	deleteOrder, args, err := sqlx.In("DELETE FROM orders WHERE id IN (?);", orderIDs)
	if err != nil {
		log.Error(err)
		return err
	}
	if _, err := tx.Exec(deleteOrder, args...); err != nil {
		log.Error(err)
		return err
	}

	deleteCart := "DELETE FROM carts WHERE id = ?;"
	if _, err := tx.Exec(deleteCart, cart.ID); err != nil {
		log.Error(err)
		return err
	}
	tx.Commit()
	log.WithFields(logrus.Fields{
		"cart": cart,
	}).Info("checkout success")
	return nil
}

func (r mysqlCartRepository) FindOrder(cartID, BookID int) (domain.Order, error) {
	var orderTable orderWithBookTable
	query := `
		SELECT o.id, cart_id, book_id, qty, title, cover, price
		FROM orders AS o
				 LEFT JOIN books b ON o.book_id = b.id
		WHERE cart_id = ?
		  AND book_id = ? LIMIT 1`

	if err := r.db.Get(&orderTable, query, cartID, BookID); err != nil {
		log.Printf("couldn't query order: %v", err)
		return domain.Order{}, err
	}
	return domain.Order{
		ID: orderTable.ID,
		Book: domain.Book{
			ID:    orderTable.BookID,
			Title: orderTable.Title,
			Cover: orderTable.Cover,
			Price: orderTable.Price,
		},
		Qty: orderTable.Qty,
	}, nil
}

func (r mysqlCartRepository) CreateNewOrder(cartID, bookID int) (domain.Order, error) {
	query := "INSERT INTO orders (cart_id, book_id, qty) VALUE (?,?,1)"
	tx := r.db.MustBegin()
	if _, err := tx.Exec(query, cartID, bookID); err != nil {
		log.Printf("couldn't insert to cart %v", err)
		tx.Rollback()
		return domain.Order{}, err
	}
	tx.Commit()
	return r.FindOrder(cartID, bookID)
}

func (r mysqlCartRepository) UpdateOrderQty(orderID, qty int) error {
	query := "UPDATE orders SET qty = ? WHERE id = ?"
	tx := r.db.MustBegin()
	if _, err := tx.Exec(query, qty, orderID); err != nil {
		log.Printf("couldn't update order %v", err)
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (r mysqlCartRepository) DeleteOrder(orderID int) error {
	query := "DELETE FROM orders WHERE id = ?;"
	tx := r.db.MustBegin()
	if _, err := tx.Exec(query, orderID); err != nil {
		log.Printf("couldn't delete order %v", err)
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (r mysqlCartRepository) FindCartByPosID(posID int) (domain.Cart, error) {
	var cartTable cartTable
	cartQuery := "SELECT id, pos_terminal_id FROM carts WHERE pos_terminal_id = ?;"
	if err := r.db.Get(&cartTable, cartQuery, posID); err != nil {
		if err != sql.ErrNoRows {
			log.Printf("couldn't query from cart: %v", err)
		}
		return domain.Cart{}, err
	}
	var cart = domain.Cart{
		ID:            cartTable.ID,
		PosTerminalID: cartTable.PosTerminalID,
	}
	orders, err := r.GetOrders(cart.ID)
	if err == sql.ErrNoRows || err == nil {
		cart.Orders = orders
		return cart, nil
	}
	return domain.Cart{}, err
}

func (r mysqlCartRepository) GetOrders(cartID int) ([]domain.Order, error) {
	var orderTable []orderWithBookTable
	query := `
		SELECT o.id, cart_id, book_id, qty, title, cover, price
		FROM orders AS o
				 LEFT JOIN books b ON o.book_id = b.id
		WHERE cart_id = ?`

	if err := r.db.Select(&orderTable, query, cartID); err != nil {
		if err != sql.ErrNoRows {
			log.Printf("couldn't query order: %v", err)
		}
		return nil, err
	}
	var orders []domain.Order
	for _, o := range orderTable {
		orders = append(orders, domain.Order{
			ID: o.ID,
			Book: domain.Book{
				ID:    o.BookID,
				Title: o.Title,
				Cover: o.Cover,
				Price: o.Price,
			},
			Qty: o.Qty,
		})
	}
	return orders, nil
}
