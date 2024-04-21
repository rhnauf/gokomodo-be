package service

import (
	"context"
	"net/http"

	"github.com/rhnauf/gokomodo-be/internal/entity"
)

type ProductRepositoryBuyer interface {
	BuyerGetListProduct(ctx context.Context) ([]*entity.Product, error)
	GetProductById(ctx context.Context, productId int) (*entity.Product, error)
}

type OrderRepositoryBuyer interface {
	InsertOrder(ctx context.Context, order *entity.Order) (*entity.Order, error)
	BuyerGetListOrder(ctx context.Context, buyerId int) ([]*entity.Order, error)
}

type UserRepositoryBuyer interface {
	GetBuyerById(ctx context.Context, buyerId int) (*entity.Buyer, error)
	GetSellerById(ctx context.Context, sellerId int) (*entity.Seller, error)
}

type buyerService struct {
	productRepo ProductRepositoryBuyer
	orderRepo   OrderRepositoryBuyer
	userRepo    UserRepositoryBuyer
}

func NewBuyerService(productRepo ProductRepositoryBuyer, orderRepo OrderRepositoryBuyer, userRepo UserRepositoryBuyer) *buyerService {
	return &buyerService{
		productRepo: productRepo,
		orderRepo:   orderRepo,
		userRepo:    userRepo,
	}
}

func (b *buyerService) GetListProduct(ctx context.Context) (*entity.ResponseDTO, error) {
	res, err := b.productRepo.BuyerGetListProduct(ctx)
	if err != nil {
		return nil, err
	}

	return entity.HandleResponseDTO(http.StatusOK, "success get list product", entity.TransformProductListToProductListResponse(res)), nil
}

func (b *buyerService) CreateOrder(ctx context.Context, request *entity.OrderRequest, buyerId int) (*entity.ResponseDTO, error) {
	order := entity.TransformOrderRequestToOrder(request)

	// should've wrapped all the operation with transaction to ensure atomicity
	// get product
	product, err := b.productRepo.GetProductById(ctx, request.ProductId)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return entity.HandleResponseDTO(http.StatusNotFound, "failed, product not found", nil), nil
	}

	// get buyer
	buyer, err := b.userRepo.GetBuyerById(ctx, buyerId)
	if err != nil {
		return nil, err
	}
	if buyer == nil {
		return entity.HandleResponseDTO(http.StatusBadRequest, "failed, buyer not found", nil), nil
	}

	// get seller
	seller, err := b.userRepo.GetSellerById(ctx, product.SellerId)
	if err != nil {
		return nil, err
	}
	if seller == nil {
		return entity.HandleResponseDTO(http.StatusBadRequest, "failed, seller not found", nil), nil
	}

	// construct
	order.SetBuyerId(buyerId)
	order.SetSellerId(seller.Id)
	order.SetAddressSource(seller.AddressPickup)
	order.SetAddressDestination(buyer.AddressSend)
	order.SetPrice(product.Price)
	order.CalculateTotalPrice()

	// create order
	res, err := b.orderRepo.InsertOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	return entity.HandleResponseDTO(http.StatusOK, "success create order", entity.TransformOrderToOrderResponse(res)), nil
}

func (b *buyerService) GetListOrder(ctx context.Context, buyerId int) (*entity.ResponseDTO, error) {
	res, err := b.orderRepo.BuyerGetListOrder(ctx, buyerId)
	if err != nil {
		return nil, err
	}

	return entity.HandleResponseDTO(http.StatusOK, "success get list order", entity.TransformOrderListToOrderListResponse(res)), nil
}
