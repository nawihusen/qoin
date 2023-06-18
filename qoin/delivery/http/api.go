package http

import (
	"Qoin/domain"
	"Qoin/qoin/delivery/http/handler"

	// "Qoin/saksi/delivery/http/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

// RouterAPI is main router for this Service Saksi REST API
func RouterAPI(app *fiber.App, qoin domain.QoinUsecase) {
	handlerQoin := &handler.QoinHandler{QoinUsecase: qoin}

	basePath := viper.GetString("server.base_path")
	path := app.Group(basePath)

	// Management
	path.Post("/order", handlerQoin.Order) //
	path.Get("/menu", handlerQoin.ListMenu)
	path.Post("/menu", handlerQoin.InsertMenu)
	path.Get("/menu/stock/:id", handlerQoin.GetStock)
	path.Get("/income", handlerQoin.GetIncome) //
	// path.Get("/detail", handlerQoin.Detail)
}
