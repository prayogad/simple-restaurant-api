package service

import (
	"context"
	"database/sql"
	"simple-restaurant-web/exceptions"
	"simple-restaurant-web/helper"
	"simple-restaurant-web/model/domain"
	"simple-restaurant-web/model/web"
	"simple-restaurant-web/repository"

	"github.com/go-playground/validator"
)

type FoodServiceImpl struct {
	FoodRepository repository.FoodRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewFoodService(foodRepository repository.FoodRepository, DB *sql.DB, validate *validator.Validate) FoodService {
	return &FoodServiceImpl{
		FoodRepository: foodRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *FoodServiceImpl) Create(ctx context.Context, request web.FoodCreateRequest) web.FoodResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	food := domain.Food{
		Name:  request.Name,
		Price: request.Price,
		Stock: request.Stock,
	}

	food = service.FoodRepository.Save(ctx, tx, food)

	return helper.ToFoodResponse(food)
}

func (service *FoodServiceImpl) Update(ctx context.Context, request web.FoodUpdateRequest) web.FoodResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_, err = service.FoodRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	food := service.FoodRepository.Update(ctx, tx, domain.Food(request))

	return helper.ToFoodResponse(food)
}

func (service *FoodServiceImpl) Delete(ctx context.Context, foodId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	food, err := service.FoodRepository.FindById(ctx, tx, foodId)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	service.FoodRepository.Delete(ctx, tx, food)
}

func (service *FoodServiceImpl) FindById(ctx context.Context, foodId int) web.FoodResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	food, err := service.FoodRepository.FindById(ctx, tx, foodId)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	return helper.ToFoodResponse(food)
}

func (service *FoodServiceImpl) FindAll(ctx context.Context) []web.FoodResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	foods := service.FoodRepository.FindAll(ctx, tx)

	return helper.ToFoodResponses(foods)
}
