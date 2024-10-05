package repository

import (
	"context"
	"database/sql"
	"simple-restaurant-web/model/domain"
)

type OrderRepository interface {
	Save(ctx context.Context, Tx *sql.Tx, order domain.Orders) domain.Orders
	Get(ctx context.Context, Tx *sql.Tx) []domain.Orders
	GetDetail(ctx context.Context, Tx *sql.Tx, orderId int) domain.Orders
}
