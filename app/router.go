package app

import (
	"simple-restaurant-web/controller"
	"simple-restaurant-web/exceptions"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(customerController controller.CustomerController, foodController controller.FoodController) *httprouter.Router {
	router := httprouter.New()

	// Customer API
	router.POST("/customer/register", customerController.Create)
	router.POST("/customer/login", customerController.Login)
	router.POST("/customer/logout/:id", customerController.Logout)
	router.PUT("/customer/update/:id", customerController.Update)
	router.DELETE("/customer/delete/:id", customerController.Delete)
	router.GET("/customer/find/:id", customerController.FindById)
	router.GET("/customer/findAll", customerController.FindAll)

	// Food API
	router.POST("/food/create", foodController.Create)

	router.PanicHandler = exceptions.ErrorHandler

	return router
}
