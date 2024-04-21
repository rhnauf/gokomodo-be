package api

import (
	"github.com/go-chi/chi/v5"

	"github.com/rhnauf/gokomodo-be/internal/controller"
	"github.com/rhnauf/gokomodo-be/internal/repository"
	"github.com/rhnauf/gokomodo-be/internal/service"
)

func (h *Handler) buyerRoutes(r chi.Router) {
	dbClient := h.DBClient.Client

	productRepository := repository.NewProductRepository(dbClient)
	orderRepository := repository.NewOrderRepository(dbClient)
	userRepository := repository.NewUserRepository(dbClient)

	buyerService := service.NewBuyerService(productRepository, orderRepository, userRepository)

	buyerController := controller.NewBuyerController(buyerService)

	r.Use(AuthJWTBuyer)
	r.Get("/list-products-buyer", buyerController.GetListProduct)
	r.Post("/orders", buyerController.CreateOrder)
	r.Get("/list-orders-buyer", buyerController.GetListOrder)
}
