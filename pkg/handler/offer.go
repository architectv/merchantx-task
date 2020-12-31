package handler

import (
	"log"
	"runtime"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func implementMe() {
	pc, fn, line, _ := runtime.Caller(1)
	log.Printf("Implement me in %s[%s:%d]\n", runtime.FuncForPC(pc).Name(), fn, line)
}

func (h *Handler) put(ctx *fiber.Ctx) error {
	implementMe()
	return ctx.SendString("PUT")
}

func (h *Handler) get(ctx *fiber.Ctx) error {
	var sellerId, offerId int
	var err error
	substr := ctx.Query("substr")
	sellerIdQuery := ctx.Query("id")
	offerIdQuery := ctx.Query("offer_id")

	if sellerIdQuery != "" {
		sellerId, err = strconv.Atoi(sellerIdQuery)
		if err != nil {
			return ctx.JSON(err)
		}
	}

	if offerIdQuery != "" {
		offerId, err = strconv.Atoi(offerIdQuery)
		if err != nil {
			return ctx.JSON(err)
		}
	}

	offers, err := h.services.Offer.Get(sellerId, offerId, substr)
	if err != nil {
		return ctx.JSON(err)
	}

	return ctx.JSON(offers)
}
