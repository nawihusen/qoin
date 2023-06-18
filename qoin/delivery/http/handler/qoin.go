package handler

import (
	"Qoin/domain"
	"Qoin/helper"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

// AccountHandler is REST API handler for Service Account System
type QoinHandler struct {
	QoinUsecase domain.QoinUsecase
}

func (qoin QoinHandler) Order(c *fiber.Ctx) (err error) {
	var input domain.OrderRequst
	err = c.BodyParser(&input)
	if err != nil {
		log.Error(err)
		fmt.Println(input)
		return helper.HTTPSimpleResponse(c, fasthttp.StatusBadRequest)
	}
	validate := validator.New()
	err = validate.Struct(input)
	if err != nil {
		if strings.Contains(err.Error(), "required") {
			return c.Status(400).SendString("Tolong isi semua bagian")
		}
		return helper.HTTPSimpleResponse(c, fasthttp.StatusBadRequest)
	}

	if len(input.Amount) != len(input.MenuID) {
		err := errors.New("invalid data")
		log.Error(err)
		return helper.HTTPSimpleResponse(c, fasthttp.StatusBadRequest)
	}

	response, err := qoin.QoinUsecase.MakeOrder(c.Context(), &input)
	if err != nil {
		log.Error(err)
		return helper.HTTPSimpleResponse(c, fasthttp.StatusInternalServerError)
	}

	return c.Status(fasthttp.StatusOK).JSON(response)
}

func (qoin QoinHandler) GetStock(c *fiber.Ctx) (err error) {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		log.Error(err)
		return helper.HTTPSimpleResponse(c, fasthttp.StatusBadRequest)
	}

	stock, err := qoin.QoinUsecase.GetStock(c.Context(), id)
	if err != nil {
		log.Error(err)
		return helper.HTTPSimpleResponse(c, fasthttp.StatusInternalServerError)
	}

	return c.Status(fasthttp.StatusOK).JSON(stock)
}

func (qoin QoinHandler) ListMenu(c *fiber.Ctx) (err error) {
	menu, err := qoin.QoinUsecase.GetListMenu(c.Context())
	if err != nil {
		log.Error(err)
		return helper.HTTPSimpleResponse(c, fasthttp.StatusInternalServerError)
	}

	return c.Status(fasthttp.StatusOK).JSON(menu)
}

func (qoin QoinHandler) InsertMenu(c *fiber.Ctx) (err error) {
	var input domain.Menu
	err = c.BodyParser(&input)
	if err != nil {
		log.Error(err)
		return helper.HTTPSimpleResponse(c, fasthttp.StatusBadRequest)
	}

	validate := validator.New()
	err = validate.Struct(input)
	if err != nil {
		if strings.Contains(err.Error(), "required") {
			return c.Status(400).SendString("Tolong isi semua bagian")
		}
		return helper.HTTPSimpleResponse(c, fasthttp.StatusBadRequest)
	}

	err = qoin.QoinUsecase.AddNewMenu(c.Context(), input)
	if err != nil {
		log.Error(err)
		return helper.HTTPSimpleResponse(c, fasthttp.StatusInternalServerError)
	}

	return c.Status(fasthttp.StatusOK).JSON("Success")
}

func (qoin QoinHandler) GetIncome(c *fiber.Ctx) (err error) {
	// format waktu yyyy-mm-dd
	from := c.Query("from")
	if from == "" {
		log.Error("from cannot blank")
		return helper.HTTPSimpleResponse(c, fasthttp.StatusBadRequest)
	}

	to := c.Query("to")
	if to == "" {
		log.Error("to cannot blank")
		return helper.HTTPSimpleResponse(c, fasthttp.StatusBadRequest)
	}

	income, err := qoin.QoinUsecase.GetIncome(c.Context(), from, to)
	if err != nil {
		log.Error(err)
		return helper.HTTPSimpleResponse(c, fasthttp.StatusInternalServerError)
	}

	return c.Status(fasthttp.StatusOK).JSON(income)
}
