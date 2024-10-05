package controller

import (
	"net/http"
	"simple-restaurant-web/helper"
	"simple-restaurant-web/model/web"
	"simple-restaurant-web/service"
	"strconv"

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

func (controller *OrderControllerImpl) Get(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	orderResponses := controller.OrderService.Get(request.Context())

	webResponse := web.WebResponse{
		Success: true,
		Message: "Successfully Get Orders Data",
		Data:    orderResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrderControllerImpl) GetDetail(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("orderId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	orderResponse := controller.OrderService.GetDetail(request.Context(), id)

	webResponse := web.WebResponse{
		Success: true,
		Message: "Successfully Get Detail Order",
		Data:    orderResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
