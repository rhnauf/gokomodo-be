package entity

import (
	"errors"
)

type BuyerLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (b BuyerLoginRequest) Validate() error {
	if b.Email == "" {
		return errors.New("email is required")
	}

	if b.Password == "" {
		return errors.New("password is required")
	}

	return nil
}

type BuyerLoginResponse struct {
	UserId int    `json:"user_id"`
	Email  string `json:"email"`
	Token  string `json:"token"`
}

func TransformBuyerToBuyerLoginResponse(buyer *Buyer, token string) *BuyerLoginResponse {
	return &BuyerLoginResponse{
		UserId: buyer.Id,
		Email:  buyer.Email,
		Token:  token,
	}
}

type SellerLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s SellerLoginRequest) Validate() error {
	if s.Email == "" {
		return errors.New("email is required")
	}

	if s.Password == "" {
		return errors.New("password is required")
	}

	return nil
}

type SellerLoginResponse struct {
	UserId int    `json:"user_id"`
	Email  string `json:"email"`
	Token  string `json:"token"`
}

func TransformSellerToSellerLoginResponse(seller *Seller, token string) *SellerLoginResponse {
	return &SellerLoginResponse{
		UserId: seller.Id,
		Email:  seller.Email,
		Token:  token,
	}
}
