package handler

import (
	"github.com/architectv/merchantx-task/pkg/service"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes(router fiber.Router) {
	api := router.Group("/api")
	{
		offers := api.Group("/offers")
		{
			offers.Put("/", h.putOffers)
			offers.Get("/", h.getOffers)
		}
		stats := api.Group("/stats")
		{
			stats.Get("/:id", h.getStat)
		}
	}
}
