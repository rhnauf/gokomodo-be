package entity

import "errors"

type Order struct {
	Id                 int     `db:"id"`
	BuyerId            int     `db:"buyer_id"`
	SellerId           int     `db:"seller_id"`
	ProductId          int     `db:"product_id"`
	AddressSource      string  `db:"address_source"`
	AddressDestination string  `db:"address_destination"`
	Qty                int     `db:"qty"`
	Price              float64 `db:"price"`
	TotalPrice         float64 `db:"total_price"`
	Status             string  `db:"status"`
}

func (o *Order) SetBuyerId(buyerId int) {
	o.BuyerId = buyerId
}

func (o *Order) SetSellerId(sellerId int) {
	o.SellerId = sellerId
}

func (o *Order) SetAddressSource(addressSource string) {
	o.AddressSource = addressSource
}

func (o *Order) SetAddressDestination(addressDestination string) {
	o.AddressDestination = addressDestination
}

func (o *Order) SetPrice(price float64) {
	o.Price = price
}

func (o *Order) CalculateTotalPrice() {
	var totalPrice float64

	// base price
	totalPrice += o.calculateItemPrice()

	// maybe add more calculation like promo/discount/shipping/etc..

	o.TotalPrice = totalPrice
}

func (o Order) calculateItemPrice() float64 {
	return o.Price * float64(o.Qty)
}

type OrderResponse struct {
	Id                 int     `json:"id"`
	BuyerId            int     `json:"buyer_id"`
	SellerId           int     `json:"seller_id"`
	ProductId          int     `json:"product_id"`
	AddressSource      string  `json:"address_source"`
	AddressDestination string  `json:"address_destination"`
	Qty                int     `json:"qty"`
	Price              float64 `json:"price"`
	TotalPrice         float64 `json:"total_price"`
	Status             string  `json:"status"`
}

type OrderRequest struct {
	ProductId int `json:"product_id"`
	Qty       int `json:"qty"`
}

func (o OrderRequest) Validate() error {
	if o.ProductId <= 0 {
		return errors.New("product id must be greater than 0")
	}

	if o.Qty <= 0 {
		return errors.New("product id must be greater than 0")
	}

	return nil
}

type OrderListResponse struct {
	Orders []*OrderResponse `json:"orders"`
}

func TransformOrderToOrderResponse(in *Order) *OrderResponse {
	return &OrderResponse{
		Id:                 in.Id,
		BuyerId:            in.BuyerId,
		SellerId:           in.SellerId,
		ProductId:          in.ProductId,
		AddressSource:      in.AddressSource,
		AddressDestination: in.AddressDestination,
		Qty:                in.Qty,
		Price:              in.Price,
		TotalPrice:         in.TotalPrice,
		Status:             in.Status,
	}
}

func TransformOrderListToOrderListResponse(in []*Order) *OrderListResponse {
	orders := make([]*OrderResponse, len(in))

	for i, e := range in {
		orders[i] = TransformOrderToOrderResponse(e)
	}

	return &OrderListResponse{
		Orders: orders,
	}
}

func TransformOrderRequestToOrder(in *OrderRequest) *Order {
	return &Order{
		ProductId: in.ProductId,
		Qty:       in.Qty,
	}
}
