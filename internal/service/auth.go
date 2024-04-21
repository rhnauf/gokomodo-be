package service

import (
	"context"
	"github.com/rhnauf/gokomodo-be/internal/helper"
	"log"
	"net/http"

	"github.com/rhnauf/gokomodo-be/internal/entity"
)

type AuthRepository interface {
	GetBuyerByEmail(ctx context.Context, email string) (*entity.Buyer, error)
	GetSellerByEmail(ctx context.Context, email string) (*entity.Seller, error)
}

type authService struct {
	authRepo AuthRepository
}

func NewAuthService(authRepo AuthRepository) *authService {
	return &authService{authRepo: authRepo}
}

func (a *authService) LoginBuyer(ctx context.Context, request *entity.BuyerLoginRequest) (*entity.ResponseDTO, error) {
	buyer, err := a.authRepo.GetBuyerByEmail(ctx, request.Email)
	if err != nil {
		return nil, err
	}

	if buyer == nil || !helper.CheckPassword(request.Password, buyer.Password) {
		return entity.HandleResponseDTO(http.StatusBadRequest, "invalid email or password", nil), nil
	}

	token, err := helper.GenerateJWT(buyer.Id, buyer.Email, helper.Buyer)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return entity.HandleResponseDTO(http.StatusOK, "success", entity.TransformBuyerToBuyerLoginResponse(buyer, token)), nil
}

func (a *authService) LoginSeller(ctx context.Context, request *entity.SellerLoginRequest) (*entity.ResponseDTO, error) {
	seller, err := a.authRepo.GetSellerByEmail(ctx, request.Email)
	if err != nil {
		return nil, err
	}

	if seller == nil || !helper.CheckPassword(request.Password, seller.Password) {
		return entity.HandleResponseDTO(http.StatusBadRequest, "invalid email or password", nil), nil
	}

	token, err := helper.GenerateJWT(seller.Id, seller.Email, helper.Seller)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return entity.HandleResponseDTO(http.StatusOK, "success", entity.TransformSellerToSellerLoginResponse(seller, token)), nil
}
