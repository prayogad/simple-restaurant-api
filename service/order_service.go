package service

import (
	"context"
	"simple-restaurant-web/model/web"
)

type OrderService interface {
	Create(ctx context.Context, request web.OrderCreateRequest) web.OrderResponse
	FindById(ctx context.Context, foodId int) web.OrderResponse
}
