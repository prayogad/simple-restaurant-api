package repository

import (
	"context"
	"database/sql"
	"simple-restaurant-web/model/domain"
)

type CustomerRepository interface {
	Save(ctx context.Context, Tx *sql.Tx, customer domain.Customer) (domain.Customer, error)
	Login(ctx context.Context, Tx *sql.Tx, username string, password string) (domain.Customer, error)
	ValidateToken(Tx *sql.Tx, token string) (idCustomer int, usernameCustomer string)
	Logout(ctx context.Context, Tx *sql.Tx, customerId int)
	Update(ctx context.Context, Tx *sql.Tx, customer domain.Customer) domain.Customer
	Delete(ctx context.Context, Tx *sql.Tx, customer domain.Customer)
	FindById(ctx context.Context, Tx *sql.Tx, customerId int) (domain.Customer, error)
	FindAll(ctx context.Context, Tx *sql.Tx) []domain.Customer
}
