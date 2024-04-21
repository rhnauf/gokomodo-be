package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/rhnauf/gokomodo-be/internal/entity"
	"github.com/rhnauf/gokomodo-be/internal/helper"
)

type BuyerService interface {
	GetListProduct(ctx context.Context) (*entity.ResponseDTO, error)
	CreateOrder(ctx context.Context, request *entity.OrderRequest, buyerId int) (*entity.ResponseDTO, error)
	GetListOrder(ctx context.Context, buyerId int) (*entity.ResponseDTO, error)
}

type buyerController struct {
	buyerService BuyerService
}

func NewBuyerController(buyerService BuyerService) *buyerController {
	return &buyerController{buyerService: buyerService}
}

func (b *buyerController) GetListProduct(w http.ResponseWriter, r *http.Request) {
	res, err := b.buyerService.GetListProduct(r.Context())
	if err != nil {
		helper.HandleInternalServerError(w)
		return
	}

	helper.HandleResponse(w, res.StatusCode, res.Message, res.Data)
}

func (b *buyerController) CreateOrder(w http.ResponseWriter, r *http.Request) {
	claims, err := helper.UnmarshallClaim(r.Context())
	if err != nil {
		log.Println(err)
		helper.HandleResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	orderRequest := new(entity.OrderRequest)
	if err := json.NewDecoder(r.Body).Decode(orderRequest); err != nil {
		log.Println(err)
		helper.HandleResponse(w, http.StatusBadRequest, "failed binding request", nil)
		return
	}

	if err := orderRequest.Validate(); err != nil {
		helper.HandleResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	res, err := b.buyerService.CreateOrder(r.Context(), orderRequest, claims.UserId)
	if err != nil {
		helper.HandleInternalServerError(w)
		return
	}

	helper.HandleResponse(w, res.StatusCode, res.Message, res.Data)
}

func (b *buyerController) GetListOrder(w http.ResponseWriter, r *http.Request) {
	claims, err := helper.UnmarshallClaim(r.Context())
	if err != nil {
		log.Println(err)
		helper.HandleResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	res, err := b.buyerService.GetListOrder(r.Context(), claims.UserId)
	if err != nil {
		helper.HandleInternalServerError(w)
		return
	}

	helper.HandleResponse(w, res.StatusCode, res.Message, res.Data)
}
