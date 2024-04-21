package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/rhnauf/gokomodo-be/internal/entity"
	"github.com/rhnauf/gokomodo-be/internal/helper"
)

type SellerService interface {
	GetListProduct(ctx context.Context, sellerId int) (*entity.ResponseDTO, error)
	CreateNewProduct(ctx context.Context, request *entity.ProductRequest, sellerId int) (*entity.ResponseDTO, error)
	AcceptOrder(ctx context.Context, orderId, sellerId int) (*entity.ResponseDTO, error)
	GetListOrder(ctx context.Context, sellerId int) (*entity.ResponseDTO, error)
}

type sellerController struct {
	sellerService SellerService
}

func NewSellerController(sellerService SellerService) *sellerController {
	return &sellerController{sellerService: sellerService}
}

func (s *sellerController) GetListProduct(w http.ResponseWriter, r *http.Request) {
	claims, err := helper.UnmarshallClaim(r.Context())
	if err != nil {
		log.Println(err)
		helper.HandleResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	res, err := s.sellerService.GetListProduct(r.Context(), claims.UserId)
	if err != nil {
		helper.HandleInternalServerError(w)
		return
	}

	helper.HandleResponse(w, res.StatusCode, res.Message, res.Data)
}

func (s *sellerController) CreateNewProduct(w http.ResponseWriter, r *http.Request) {
	claims, err := helper.UnmarshallClaim(r.Context())
	if err != nil {
		log.Println(err)
		helper.HandleResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	productRequest := new(entity.ProductRequest)
	if err := json.NewDecoder(r.Body).Decode(productRequest); err != nil {
		log.Println(err)
		helper.HandleResponse(w, http.StatusBadRequest, "failed binding request", nil)
		return
	}

	if err := productRequest.Validate(); err != nil {
		helper.HandleResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	res, err := s.sellerService.CreateNewProduct(r.Context(), productRequest, claims.UserId)
	if err != nil {
		helper.HandleInternalServerError(w)
		return
	}

	helper.HandleResponse(w, res.StatusCode, res.Message, res.Data)
}

func (s *sellerController) AcceptOrder(w http.ResponseWriter, r *http.Request) {
	params := chi.URLParam(r, "id")

	orderId, err := strconv.Atoi(params)
	if err != nil {
		log.Println(err)
		helper.HandleResponse(w, http.StatusBadRequest, "order id must be numeric", nil)
		return
	}

	if orderId <= 0 {
		log.Println(err)
		helper.HandleResponse(w, http.StatusBadRequest, "order id must be greater than 0", nil)
		return
	}

	claims, err := helper.UnmarshallClaim(r.Context())
	if err != nil {
		log.Println(err)
		helper.HandleResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	res, err := s.sellerService.AcceptOrder(r.Context(), orderId, claims.UserId)
	if err != nil {
		helper.HandleInternalServerError(w)
		return
	}

	helper.HandleResponse(w, res.StatusCode, res.Message, res.Data)
}

func (s *sellerController) GetListOrder(w http.ResponseWriter, r *http.Request) {
	claims, err := helper.UnmarshallClaim(r.Context())
	if err != nil {
		log.Println(err)
		helper.HandleResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	res, err := s.sellerService.GetListOrder(r.Context(), claims.UserId)
	if err != nil {
		helper.HandleInternalServerError(w)
		return
	}

	helper.HandleResponse(w, res.StatusCode, res.Message, res.Data)
}
