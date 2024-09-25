package repository

import (
	"context"
	"database/sql"
	"simple-restaurant-web/model/domain"
)

type OrderRepository interface {
	Save(ctx context.Context, Tx *sql.Tx, order domain.Orders, orderDetail []domain.OrderDetail) (domain.Orders, domain.OrderDetail)
	FindById(ctx context.Context, Tx *sql.Tx, orderId int) (domain.Orders, domain.OrderDetail, error)
	FindAll(ctx context.Context, Tx *sql.Tx) []domain.Orders
}
