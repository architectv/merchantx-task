package handler

import (
	"errors"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func (h *Handler) putOffers(ctx *fiber.Ctx) error {
	start := time.Now()
	type Input struct {
		Id   int    `json:"id"`
		Link string `json:"link"`
	}
	input := &Input{}
	if err := ctx.BodyParser(input); err != nil {
		return ctx.JSON(err)
	}
	if input.Id <= 0 || input.Link == "" {
		return ctx.JSON(errors.New("wrong input"))
	}

	// TODO: file storage
	filename := "./tmp.xlsx"
	err := downloadFile(input.Link, filename)
	if err != nil {
		return ctx.JSON(err)
	}
	duration := time.Since(start)
	log.Println("DOWNLOAD:", duration)

	start = time.Now()
	stat, err := h.services.Offer.Put(input.Id, filename)
	if err != nil {
		return ctx.JSON(err)
	}
	duration = time.Since(start)
	log.Println("PUT:", duration)

	return ctx.JSON(stat)
}

func downloadFile(url, filepath string) error {
	start := time.Now()
	var dst []byte
	_, bodyBytes, err := fasthttp.Get(dst, url)

	duration := time.Since(start)
	log.Println("LINK:", duration)

	start = time.Now()
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, strings.NewReader(string(bodyBytes)))
	duration = time.Since(start)
	log.Println("FILE:", duration)
	return err
}

func (h *Handler) getOffers(ctx *fiber.Ctx) error {
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
