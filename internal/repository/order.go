package repository

import (
	"context"
	"errors"
	"log"

	"github.com/jmoiron/sqlx"

	"github.com/rhnauf/gokomodo-be/internal/entity"
)

var (
	pending          = "pending"
	accepted         = "accepted"
	ErrOrderNotFound = errors.New("no order found")
)

type orderRepository struct {
	db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) *orderRepository {
	return &orderRepository{db: db}
}

func (o *orderRepository) InsertOrder(ctx context.Context, order *entity.Order) (*entity.Order, error) {
	res := entity.Order{}
	err := o.db.
		QueryRow(`
		INSERT INTO orders (buyer_id, seller_id, product_id, address_source, address_destination, qty, price, total_price)
		VALUES
		($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING *
	`, order.BuyerId, order.SellerId, order.ProductId, order.AddressSource, order.AddressDestination, order.Qty, order.Price, order.TotalPrice).
		Scan(&res.Id, &res.BuyerId, &res.SellerId, &res.AddressSource, &res.AddressDestination, &res.ProductId, &res.Qty, &res.Price, &res.TotalPrice, &res.Status)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &res, nil
}

func (o *orderRepository) UpdateOrder(ctx context.Context, orderId, sellerId int) error {
	res, err := o.db.Exec(`
		UPDATE orders
		SET status = $1
		WHERE id = $2 AND seller_id = $3
	`, accepted, orderId, sellerId)
	if err != nil {
		log.Println(err)
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
		return err
	}

	if rows == 0 {
		log.Printf("no order found with id %d and seller id %d\n", orderId, sellerId)
		return ErrOrderNotFound
	}

	return nil
}

func (o *orderRepository) SellerGetListOrder(ctx context.Context, sellerId int) ([]*entity.Order, error) {
	res := make([]*entity.Order, 0)
	err := o.db.Select(&res, `
		SELECT * FROM orders WHERE seller_id = $1
	`, sellerId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return res, err
}

func (o *orderRepository) BuyerGetListOrder(ctx context.Context, buyerId int) ([]*entity.Order, error) {
	res := make([]*entity.Order, 0)
	err := o.db.Select(&res, `
		SELECT * FROM orders WHERE buyer_id = $1
	`, buyerId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return res, err
}
