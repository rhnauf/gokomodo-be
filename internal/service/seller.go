package service

import (
	"context"
	"errors"
	"github.com/rhnauf/gokomodo-be/internal/repository"
	"net/http"

	"github.com/rhnauf/gokomodo-be/internal/entity"
)

type ProductRepositorySeller interface {
	SellerGetListProduct(ctx context.Context, sellerId int) ([]*entity.Product, error)
	InsertProduct(ctx context.Context, product *entity.Product) (*entity.Product, error)
}

type OrderRepositorySeller interface {
	UpdateOrder(ctx context.Context, orderId, sellerId int) error
	SellerGetListOrder(ctx context.Context, sellerId int) ([]*entity.Order, error)
}

type sellerService struct {
	productRepo ProductRepositorySeller
	orderRepo   OrderRepositorySeller
}

func NewSellerService(productRepo ProductRepositorySeller, orderRepo OrderRepositorySeller) *sellerService {
	return &sellerService{
		productRepo: productRepo,
		orderRepo:   orderRepo,
	}
}

func (s *sellerService) GetListProduct(ctx context.Context, sellerId int) (*entity.ResponseDTO, error) {
	res, err := s.productRepo.SellerGetListProduct(ctx, sellerId)
	if err != nil {
		return nil, err
	}

	return entity.HandleResponseDTO(http.StatusOK, "success get list product", entity.TransformProductListToProductListResponse(res)), nil
}

func (s *sellerService) CreateNewProduct(ctx context.Context, request *entity.ProductRequest, sellerId int) (*entity.ResponseDTO, error) {
	product := entity.TransformProductRequestToProduct(request)
	product.SetSellerId(sellerId)

	res, err := s.productRepo.InsertProduct(ctx, product)
	if err != nil {
		return nil, err
	}

	return entity.HandleResponseDTO(http.StatusCreated, "success create product", entity.TransformProductToProductResponse(res)), nil
}

func (s *sellerService) AcceptOrder(ctx context.Context, orderId, sellerId int) (*entity.ResponseDTO, error) {
	err := s.orderRepo.UpdateOrder(ctx, orderId, sellerId)
	if err != nil {
		if errors.Is(err, repository.ErrOrderNotFound) {
			return entity.HandleResponseDTO(http.StatusNotFound, "failed, order not found", nil), nil
		}
		return nil, err
	}

	return entity.HandleResponseDTO(http.StatusOK, "success accepting order", nil), nil
}

func (s *sellerService) GetListOrder(ctx context.Context, sellerId int) (*entity.ResponseDTO, error) {
	res, err := s.orderRepo.SellerGetListOrder(ctx, sellerId)
	if err != nil {
		return nil, err
	}

	return entity.HandleResponseDTO(http.StatusOK, "success get list order", entity.TransformOrderListToOrderListResponse(res)), nil
}
