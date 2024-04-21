package entity

import "errors"

type Product struct {
	Id          int     `db:"id"`
	SellerId    int     `db:"seller_id"`
	ProductName string  `db:"product_name"`
	Description string  `db:"description"`
	Price       float64 `db:"price"`
}

func (p *Product) SetSellerId(sellerId int) {
	p.SellerId = sellerId
}

type ProductRequest struct {
	ProductName string  `json:"product_name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func (p *ProductRequest) Validate() error {
	if p.ProductName == "" {
		return errors.New("product name is required")
	}

	if p.Price < 0 {
		return errors.New("price must be greater than 0")
	}

	return nil
}

func TransformProductRequestToProduct(in *ProductRequest) *Product {
	return &Product{
		ProductName: in.ProductName,
		Description: in.Description,
		Price:       in.Price,
	}
}

type ProductResponse struct {
	Id          int     `json:"id"`
	SellerId    int     `json:"seller_id"`
	ProductName string  `json:"product_name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type ProductListResponse struct {
	Products []*ProductResponse `json:"products"`
}

func TransformProductToProductResponse(in *Product) *ProductResponse {
	return &ProductResponse{
		Id:          in.Id,
		SellerId:    in.SellerId,
		ProductName: in.ProductName,
		Description: in.Description,
		Price:       in.Price,
	}
}

func TransformProductListToProductListResponse(in []*Product) *ProductListResponse {
	products := make([]*ProductResponse, len(in))

	for i, e := range in {
		products[i] = TransformProductToProductResponse(e)
	}

	return &ProductListResponse{
		Products: products,
	}
}
