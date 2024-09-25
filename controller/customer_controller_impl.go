package controller

import (
	"net/http"
	"simple-restaurant-web/helper"
	"simple-restaurant-web/model/web"
	"simple-restaurant-web/service"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

type CustomerControllerImpl struct {
	CustomerService service.CustomerService
}

func NewCustomerController(customerService service.CustomerService) CustomerController {
	return &CustomerControllerImpl{
		CustomerService: customerService,
	}
}

func (controller *CustomerControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	customerCreateRequest := web.CustomerCreateRequest{}
	helper.ReadFromRequestBody(request, &customerCreateRequest)

	customerResponse := controller.CustomerService.Create(request.Context(), customerCreateRequest)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Successfully Create Customer Data",
		Data:    customerResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CustomerControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	customerLoginRequest := web.CustomerLoginRequest{}
	helper.ReadFromRequestBody(request, &customerLoginRequest)

	customerResponse := controller.CustomerService.Login(request.Context(), customerLoginRequest)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Login Successfully",
		Data:    customerResponse,
	}

	// cookie := new(http.Cookie)
	// cookie.Name = "auth"
	// cookie.Value = customerResponse.Token
	// cookie.Path = "/"
	// http.SetCookie(writer, cookie)

	http.SetCookie(writer, &http.Cookie{
		Name:     "auth",
		Value:    customerResponse.Token,
		Path:     "/customer",
		Expires:  time.Now().Add(24 * time.Hour), // Cookie expires in 24 hours
		HttpOnly: true,                           // Make the cookie HTTP only
	})

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CustomerControllerImpl) Logout(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	customerId := params.ByName("id")
	id, err := strconv.Atoi(customerId)
	helper.PanicIfError(err)

	controller.CustomerService.Logout(request.Context(), id)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Logout Successfully",
	}

	http.SetCookie(writer, &http.Cookie{
		Name:     "auth",
		Value:    "",
		Expires:  time.Unix(0, 0), // Set expiration time to the past
		MaxAge:   -1,              // Set MaxAge to -1 to delete the cookie
		HttpOnly: true,
	})

	helper.WriteToResponseBody(writer, webResponse)

}

func (controller *CustomerControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	customerUpdateRequest := web.CustomerUpdateRequest{}
	helper.ReadFromRequestBody(request, &customerUpdateRequest)

	customerId := params.ByName("id")
	id, err := strconv.Atoi(customerId)
	helper.PanicIfError(err)
	customerUpdateRequest.Id = id

	customerResponse := controller.CustomerService.Update(request.Context(), customerUpdateRequest)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Successfully Update Customer Data",
		Data:    customerResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CustomerControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	customerId := params.ByName("id")
	id, err := strconv.Atoi(customerId)
	helper.PanicIfError(err)

	controller.CustomerService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Successfully Delete Customer Data",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CustomerControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	customerId := params.ByName("id")
	id, err := strconv.Atoi(customerId)
	helper.PanicIfError(err)

	customerResponse := controller.CustomerService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Success: true,
		Message: "Successfully Fetch Customer Data",
		Data:    customerResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CustomerControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	customerResponses := controller.CustomerService.FindAll(request.Context())
	webResponses := web.WebResponse{
		Success: true,
		Message: "Successfully Fetch All Customer Data",
		Data:    customerResponses,
	}

	helper.WriteToResponseBody(writer, webResponses)
}
