package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/rhnauf/gokomodo-be/internal/service"

	"github.com/rhnauf/gokomodo-be/internal/controller"
	"github.com/rhnauf/gokomodo-be/internal/repository"
)

func (h *Handler) authRoutes(r chi.Router) {
	authRepository := repository.NewAuthRepository(h.DBClient.Client)

	authService := service.NewAuthService(authRepository)

	authController := controller.NewAuthController(authService)

	r.Post("/login-buyer", authController.LoginBuyer)
	r.Post("/login-seller", authController.LoginSeller)
}
