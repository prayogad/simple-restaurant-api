package main

import (
	"net/http"
	"simple-restaurant-web/app"
	"simple-restaurant-web/controller"
	"simple-restaurant-web/helper"
	"simple-restaurant-web/middleware"
	"simple-restaurant-web/repository"
	"simple-restaurant-web/service"

	"github.com/go-playground/validator"
	_ "github.com/lib/pq"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	customerRepository := repository.NewCustomerRepository()
	customerService := service.NewCustomerService(customerRepository, db, validate)
	customerController := controller.NewCustomerController(customerService)

	foodRepository := repository.NewFoodRepository()
	foodService := service.NewFoodService(foodRepository, db, validate)
	foodController := controller.NewFoodController(foodService)

	router := app.NewRouter(customerController, foodController)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: middleware.NewAuthMiddleware(router, customerService),
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
