package repository

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/jmoiron/sqlx"

	"github.com/rhnauf/gokomodo-be/internal/entity"
)

type authRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *authRepository {
	return &authRepository{db: db}
}

func (a *authRepository) GetBuyerByEmail(ctx context.Context, email string) (*entity.Buyer, error) {
	buyer := &entity.Buyer{}
	err := a.db.Get(buyer, "SELECT * FROM buyers WHERE email = $1", email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("no buyer found with email %s\n", email)
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}

	return buyer, nil
}

func (a *authRepository) GetSellerByEmail(ctx context.Context, email string) (*entity.Seller, error) {
	seller := &entity.Seller{}
	err := a.db.Get(seller, "SELECT * FROM sellers WHERE email = $1", email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("no seller found with email %s\n", email)
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}

	return seller, nil
}
