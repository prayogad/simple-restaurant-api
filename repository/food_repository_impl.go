package repository

import (
	"context"
	"database/sql"
	"errors"
	"simple-restaurant-web/helper"
	"simple-restaurant-web/model/domain"
)

type FoodRepositoryImpl struct {
}

func NewFoodRepository() FoodRepository {
	return &FoodRepositoryImpl{}
}

func (repository *FoodRepositoryImpl) Save(ctx context.Context, Tx *sql.Tx, food domain.Food) domain.Food {
	var id int
	SQL := "INSERT INTO food(name, price, stock) VALUES($1, $2, $3) RETURNING id"
	Tx.QueryRowContext(ctx, SQL, food.Name, food.Price, food.Stock).Scan(&id)
	food.Id = id

	return food
}

func (repository *FoodRepositoryImpl) Update(ctx context.Context, Tx *sql.Tx, food domain.Food) domain.Food {
	currentFood := domain.Food{}
	SQL := "SELECT name, price, stock FROM food WHERE id = $1"
	err := Tx.QueryRowContext(ctx, SQL, food.Id).Scan(&currentFood.Name, &currentFood.Price, &currentFood.Stock)
	helper.PanicIfError(err)

	if food.Name != "" {
		currentFood.Name = food.Name
	}

	if food.Price != 0 {
		currentFood.Price = food.Price
	}

	if food.Stock != 0 {
		currentFood.Stock = food.Stock
	}

	SQL = "UPDATE food SET name = $1, price = $2, stock = $3 WHERE id = $4"
	_, err = Tx.ExecContext(ctx, SQL, currentFood.Name, currentFood.Price, currentFood.Stock, food.Id)
	helper.PanicIfError(err)

	return currentFood
}

func (repository *FoodRepositoryImpl) Delete(ctx context.Context, Tx *sql.Tx, food domain.Food) {
	SQL := "DELETE FROM food WHERE id = $1"
	_, err := Tx.ExecContext(ctx, SQL, food.Id)
	helper.PanicIfError(err)
}

func (repository *FoodRepositoryImpl) FindById(ctx context.Context, Tx *sql.Tx, foodId int) (domain.Food, error) {
	SQL := "SELECT id, name, price, stock FROM food WHERE id = $1"
	rows, err := Tx.QueryContext(ctx, SQL, foodId)
	helper.PanicIfError(err)
	defer rows.Close()

	food := domain.Food{}
	if rows.Next() {
		err := rows.Scan(&food.Id, &food.Name, &food.Price, &food.Stock)
		helper.PanicIfError(err)
		return food, nil
	} else {
		return food, errors.New("food data not found")
	}
}

func (repository *FoodRepositoryImpl) FindAll(ctx context.Context, Tx *sql.Tx) []domain.Food {
	SQL := "SELECT id, name, price, stock FROM food"
	rows, err := Tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var foods []domain.Food
	for rows.Next() {
		food := domain.Food{}
		err := rows.Scan(&food.Id, &food.Name, &food.Price, &food.Stock)
		helper.PanicIfError(err)
		foods = append(foods, food)
	}

	return foods
}
