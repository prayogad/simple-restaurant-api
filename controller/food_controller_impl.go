package controller

import (
	"net/http"
	"simple-restaurant-web/helper"
	"simple-restaurant-web/model/web"
	"simple-restaurant-web/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type FoodControllerImpl struct {
	FoodService service.FoodService
}

func NewFoodController(foodService service.FoodService) FoodController {
	return &FoodControllerImpl{
		FoodService: foodService,
	}
}

func (controller *FoodControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	foodCreateRequest := web.FoodCreateRequest{}
	helper.ReadFromRequestBody(request, &foodCreateRequest)

	foodResponse := controller.FoodService.Create(request.Context(), foodCreateRequest)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Successfully Create Food Data",
		Data:    foodResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *FoodControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	foodUpdateRequest := web.FoodUpdateRequest{}
	helper.ReadFromRequestBody(request, &foodUpdateRequest)

	foodId := params.ByName("foodId")
	id, err := strconv.Atoi(foodId)
	helper.PanicIfError(err)
	foodUpdateRequest.Id = id

	foodResponse := controller.FoodService.Update(request.Context(), foodUpdateRequest)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Successfully Update Food Data",
		Data:    foodResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *FoodControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	foodId := params.ByName("foodId")
	id, err := strconv.Atoi(foodId)
	helper.PanicIfError(err)

	controller.FoodService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Successfully Delete Food Data",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *FoodControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	foodId := params.ByName("foodId")
	id, err := strconv.Atoi(foodId)
	helper.PanicIfError(err)

	foodResponse := controller.FoodService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Successfully Create Food Data",
		Data:    foodResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *FoodControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	foodResponses := controller.FoodService.FindAll(request.Context())
	webResponses := web.WebResponse{
		Success: true,
		Message: "Successfully Create Food Data",
		Data:    foodResponses,
	}

	helper.WriteToResponseBody(writer, webResponses)
}
