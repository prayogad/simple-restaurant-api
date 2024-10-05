package app

import (
	"simple-restaurant-web/controller"
	"simple-restaurant-web/exceptions"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(customerController controller.CustomerController, foodController controller.FoodController, orderController controller.OrderController) *httprouter.Router {
	router := httprouter.New()

	// Customer API
	router.POST("/customer/register", customerController.Create)
	router.POST("/customer/login", customerController.Login)
	router.POST("/customer/logout", customerController.Logout)
	router.PUT("/customer/update", customerController.Update)
	router.DELETE("/customer/delete/:id", customerController.Delete)
	router.GET("/customer/current", customerController.FindById)
	router.GET("/customer/findAll", customerController.FindAll)

	// Food API
	router.POST("/food/create", foodController.Create)
	router.PUT("/food/update/:foodId", foodController.Update)
	router.DELETE("/food/delete/:foodId", foodController.Delete)
	router.GET("/food/:foodId", foodController.FindById)
	router.GET("/food", foodController.FindAll)

	// Order
	router.POST("/customer/order", orderController.Create)
	router.GET("/customer/order", orderController.Get)
	router.GET("/customer/orderDetail/:orderId", orderController.GetDetail)

	router.PanicHandler = exceptions.ErrorHandler

	return router
}
