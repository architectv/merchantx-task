package handler

import (
	"fmt"
	"runtime"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
}

func (h *Handler) InitRoutes(router fiber.Router) {
	api := router.Group("/api")
	{
		offers := api.Group("/offers")
		{
			offers.Put("/", h.Put)
			offers.Get("/", h.Get)
		}
	}
}

func implementMe() {
	pc, fn, line, _ := runtime.Caller(1)
	fmt.Printf("Implement me in %s[%s:%d]\n", runtime.FuncForPC(pc).Name(), fn, line)
}

func (h *Handler) Put(ctx *fiber.Ctx) error {
	implementMe()
	return ctx.SendString("PUT")
}

func (h *Handler) Get(ctx *fiber.Ctx) error {
	implementMe()
	return ctx.SendString("GET")
}
