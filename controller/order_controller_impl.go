package controller

import (
	"net/http"
	"simple-restaurant-web/helper"
	"simple-restaurant-web/model/web"
	"simple-restaurant-web/service"

	"github.com/julienschmidt/httprouter"
)

type OrderControllerImpl struct {
	OrderService service.OrderService
}

func NewOrderController(orderService service.OrderService) OrderController {
	return &OrderControllerImpl{
		OrderService: orderService,
	}
}

func (controller *OrderControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	orderCreateRequest := web.OrderCreateRequest{}
	helper.ReadFromRequestBody(request, &orderCreateRequest)

	orderResponse := controller.OrderService.Create(request.Context(), orderCreateRequest)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Successfully Create Order Data",
		Data:    orderResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
