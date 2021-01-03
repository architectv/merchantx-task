package handler

import (
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

const (
	XlsxDir = "./data/"
)

func (h *Handler) getStat(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil || id == 0 {
		ctx.Status(fiber.StatusBadRequest)
		return ctx.SendString(err.Error())
	}

	stat, err := h.services.Offer.GetStat(id)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return ctx.SendString(err.Error())
	}

	return ctx.JSON(stat)
}

func getXlsxFilename(id int) string {
	filename := XlsxDir + "id" + strconv.Itoa(id) +
		"_" + time.Now().Format("2006-01-02_15-04-05.000000") + ".xlsx"
	return filename
}

func (h *Handler) putOffers(ctx *fiber.Ctx) error {
	type Input struct {
		Id   int    `json:"id"`
		Link string `json:"link"`
	}
	input := &Input{}
	if err := ctx.BodyParser(input); err != nil {
		ctx.Status(fiber.StatusBadRequest)
		return ctx.SendString(err.Error())
	}
	if input.Id <= 0 || input.Link == "" {
		ctx.Status(fiber.StatusBadRequest)
		return ctx.SendString("wrong input")
	}

	statId, err := h.services.Offer.CreateStat()
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return ctx.SendString(err.Error())
	}

	go func() {
		filename := getXlsxFilename(input.Id)
		if err := downloadFile(input.Link, filename); err != nil {
			h.services.Offer.ErrorStat(statId, "wrong link")
			return
		}

		h.services.Offer.PutWithFile(input.Id, statId, filename)
	}()

	return ctx.JSON(fiber.Map{"stat_id": statId})
}

func downloadFile(url, filepath string) error {
	var dst []byte
	status, bodyBytes, err := fasthttp.Get(dst, url)
	if status != fiber.StatusOK {
		return errors.New("error while get file")
	}

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, strings.NewReader(string(bodyBytes)))
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
			ctx.Status(fiber.StatusBadRequest)
			return ctx.SendString(err.Error())
		}
	}

	if offerIdQuery != "" {
		offerId, err = strconv.Atoi(offerIdQuery)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest)
			return ctx.SendString(err.Error())
		}
	}

	offers, err := h.services.Offer.GetAllByParams(sellerId, offerId, substr)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return ctx.SendString(err.Error())
	}

	return ctx.JSON(offers)
}
