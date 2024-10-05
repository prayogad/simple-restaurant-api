package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type OrderController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Get(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetDetail(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
