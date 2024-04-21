package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/jmoiron/sqlx"

	"github.com/rhnauf/gokomodo-be/internal/entity"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *userRepository {
	return &userRepository{db: db}
}

func (u *userRepository) GetBuyerById(ctx context.Context, buyerId int) (*entity.Buyer, error) {
	buyer := &entity.Buyer{}
	err := u.db.Get(buyer, "SELECT * FROM buyers WHERE id = $1", buyerId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("no buyer found with id %d\n", buyerId)
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}

	return buyer, nil
}

func (u *userRepository) GetSellerById(ctx context.Context, sellerId int) (*entity.Seller, error) {
	seller := &entity.Seller{}
	err := u.db.Get(seller, "SELECT * FROM sellers WHERE id = $1", sellerId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("no seller found with id %d\n", sellerId)
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}

	return seller, nil
}
