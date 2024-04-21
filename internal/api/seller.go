package api

import (
	"github.com/go-chi/chi/v5"

	"github.com/rhnauf/gokomodo-be/internal/controller"
	"github.com/rhnauf/gokomodo-be/internal/repository"
	"github.com/rhnauf/gokomodo-be/internal/service"
)

func (h *Handler) sellerRoutes(r chi.Router) {
	dbClient := h.DBClient.Client

	productRepository := repository.NewProductRepository(dbClient)
	orderRepository := repository.NewOrderRepository(dbClient)

	sellerService := service.NewSellerService(productRepository, orderRepository)

	sellerController := controller.NewSellerController(sellerService)

	r.Use(AuthJWTSeller)
	r.Get("/list-products-seller", sellerController.GetListProduct)
	r.Post("/products", sellerController.CreateNewProduct)
	r.Get("/list-orders-seller", sellerController.GetListOrder)
	r.Put("/orders/{id}", sellerController.AcceptOrder)
}
