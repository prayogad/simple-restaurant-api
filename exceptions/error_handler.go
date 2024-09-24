package exceptions

import (
	"fmt"
	"net/http"
	"simple-restaurant-web/helper"
	"simple-restaurant-web/model/web"

	"github.com/go-playground/validator"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {

	if notFoundError(writer, request, err) {
		return
	}

	if usernameTakenError(writer, request, err) {
		return
	}

	if validationError(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func validationError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exceptions, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Success: false,
			Message: "Bad Request",
			Data:    exceptions.Error(),
		}

		helper.WriteToResponseBody(writer, webResponse)

		return true
	} else {
		return false
	}
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exceptions, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Success: false,
			Message: "Not Found",
			Data:    exceptions.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func usernameTakenError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	_, ok := err.(UsernameTakenError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusConflict)

		webResponse := web.WebResponse{
			Success: false,
			Message: "Username Already Exist",
		}

		helper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}

}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Success: false,
		Message: "Internal Server Error",

		//disini kita convert ke string, karena JSON encode tidak bisa konversi error bawaaan Go secara langsung ke JSON
		Data: fmt.Sprintf("%v", err),
	}

	helper.WriteToResponseBody(writer, webResponse)
}
