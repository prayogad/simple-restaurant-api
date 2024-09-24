package repository

import (
	"context"
	"database/sql"
	"simple-restaurant-web/model/domain"
)

type FoodRepository interface {
	Save(ctx context.Context, Tx *sql.Tx, food domain.Food) domain.Food
	Update(ctx context.Context, Tx *sql.Tx, food domain.Food) domain.Food
	Delete(ctx context.Context, Tx *sql.Tx, food domain.Food)
	FindById(ctx context.Context, Tx *sql.Tx, foodId int) (domain.Food, error)
	FindAll(ctx context.Context, Tx *sql.Tx) []domain.Food
}
