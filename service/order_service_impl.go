package service

import (
	"context"
	"database/sql"
	"simple-restaurant-web/helper"
	"simple-restaurant-web/model/domain"
	"simple-restaurant-web/model/web"
	"simple-restaurant-web/repository"

	"github.com/go-playground/validator"
)

type OrderServiceImpl struct {
	OrderRepository repository.OrderRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewOrderService(orderRepository repository.OrderRepository, DB *sql.DB, validate *validator.Validate) OrderService {
	return &OrderServiceImpl{
		OrderRepository: orderRepository,
		DB:              DB,
		Validate:        validate,
	}
}

func (service *OrderServiceImpl) Create(ctx context.Context, request web.OrderCreateRequest) web.OrderResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	orders := domain.Orders{}

	for _, orderDetail := range request.OrderDetails {
		newOrderDetail := domain.OrderDetail{
			FoodId:   orderDetail.FoodId,
			Quantity: orderDetail.Quantity,
		}
		orders.OrderDetails = append(orders.OrderDetails, newOrderDetail)
	}

	orderResponse := service.OrderRepository.Save(ctx, tx, orders)

	return helper.ToOrderResponse(orderResponse)
}

func (service *OrderServiceImpl) Get(ctx context.Context) []web.OrderResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	orderResponse := service.OrderRepository.Get(ctx, tx)

	return helper.ToOrderResponses(orderResponse)
}

func (service *OrderServiceImpl) GetDetail(ctx context.Context, orderId int) web.OrderResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	orderResponse := service.OrderRepository.GetDetail(ctx, tx, orderId)

	return helper.ToOrderResponse(orderResponse)
}
