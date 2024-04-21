package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/rhnauf/gokomodo-be/internal/entity"
	"github.com/rhnauf/gokomodo-be/internal/helper"
)

type AuthService interface {
	LoginBuyer(ctx context.Context, request *entity.BuyerLoginRequest) (*entity.ResponseDTO, error)
	LoginSeller(ctx context.Context, request *entity.SellerLoginRequest) (*entity.ResponseDTO, error)
}

type authController struct {
	authService AuthService
}

func NewAuthController(authService AuthService) *authController {
	return &authController{authService: authService}
}

func (a *authController) LoginBuyer(w http.ResponseWriter, r *http.Request) {
	buyerLogin := new(entity.BuyerLoginRequest)

	if err := json.NewDecoder(r.Body).Decode(buyerLogin); err != nil {
		log.Println(err)
		helper.HandleResponse(w, http.StatusBadRequest, "failed binding request", nil)
		return
	}

	if err := buyerLogin.Validate(); err != nil {
		log.Println(err)
		helper.HandleResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	res, err := a.authService.LoginBuyer(r.Context(), buyerLogin)
	if err != nil {
		helper.HandleInternalServerError(w)
		return
	}

	helper.HandleResponse(w, res.StatusCode, res.Message, res.Data)
}

func (a *authController) LoginSeller(w http.ResponseWriter, r *http.Request) {
	sellerLogin := new(entity.SellerLoginRequest)

	if err := json.NewDecoder(r.Body).Decode(sellerLogin); err != nil {
		log.Println(err)
		helper.HandleResponse(w, http.StatusBadRequest, "failed binding request", nil)
		return
	}

	if err := sellerLogin.Validate(); err != nil {
		log.Println(err)
		helper.HandleResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	res, err := a.authService.LoginSeller(r.Context(), sellerLogin)
	if err != nil {
		helper.HandleInternalServerError(w)
		return
	}

	helper.HandleResponse(w, res.StatusCode, res.Message, res.Data)
}
