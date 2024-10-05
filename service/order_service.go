package service

import (
	"context"
	"simple-restaurant-web/model/web"
)

type OrderService interface {
	Create(ctx context.Context, request web.OrderCreateRequest) web.OrderResponse
	Get(ctx context.Context) []web.OrderResponse
	GetDetail(ctx context.Context, orderId int) web.OrderResponse
}
