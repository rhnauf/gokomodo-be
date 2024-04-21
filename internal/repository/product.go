package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/jmoiron/sqlx"

	"github.com/rhnauf/gokomodo-be/internal/entity"
)

type productRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *productRepository {
	return &productRepository{db: db}
}

func (p *productRepository) InsertProduct(ctx context.Context, product *entity.Product) (*entity.Product, error) {
	res := entity.Product{}
	err := p.db.QueryRow(`
		INSERT INTO products (seller_id, product_name, description, price)
		VALUES
		($1, $2, $3, $4)
		RETURNING id, seller_id, product_name, description, price
	`, product.SellerId, product.ProductName, product.Description, product.Price).Scan(&res.Id, &res.SellerId, &res.ProductName, &res.Description, &res.Price)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &res, nil
}

func (p *productRepository) SellerGetListProduct(ctx context.Context, sellerId int) ([]*entity.Product, error) {
	rows, err := p.db.Query(`
		SELECT *
		FROM products
		WHERE seller_id = $1
	`, sellerId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	res := make([]*entity.Product, 0)

	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(&product.Id, &product.ProductName, &product.Description, &product.Price, &product.SellerId); err != nil {
			log.Println("error scanning rows =>", err)
			return nil, err
		}
		res = append(res, &product)
	}

	return res, nil
}

func (p *productRepository) BuyerGetListProduct(ctx context.Context) ([]*entity.Product, error) {
	rows, err := p.db.Query(`
		SELECT *
		FROM products
	`)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	res := make([]*entity.Product, 0)

	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(&product.Id, &product.ProductName, &product.Description, &product.Price, &product.SellerId); err != nil {
			log.Println("error scanning rows =>", err)
			return nil, err
		}
		res = append(res, &product)
	}

	return res, nil
}

func (p *productRepository) GetProductById(ctx context.Context, productId int) (*entity.Product, error) {
	res := entity.Product{}
	err := p.db.QueryRow(`
		SELECT *
		FROM products
		WHERE id = $1
	`, productId).Scan(&res.Id, &res.ProductName, &res.Description, &res.Price, &res.SellerId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("no product found with id %d\n", productId)
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}

	return &res, nil
}
