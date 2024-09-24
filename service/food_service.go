package service

import (
	"context"
	"simple-restaurant-web/model/web"
)

type FoodService interface {
	Create(ctx context.Context, request web.FoodCreateRequest) web.FoodResponse
	Update(ctx context.Context, request web.FoodUpdateRequest) web.FoodResponse
	Delete(ctx context.Context, foodId int)
	FindById(ctx context.Context, foodId int) web.FoodResponse
	FindAll(ctx context.Context) []web.FoodResponse
}
