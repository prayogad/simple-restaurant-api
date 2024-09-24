package service

import (
	"context"
	"simple-restaurant-web/model/web"
)

type CustomerService interface {
	Create(ctx context.Context, request web.CustomerCreateRequest) web.CustomerResponse
	Login(ctx context.Context, request web.CustomerLoginRequest) web.CustomerResponse
	ValidateToken(token string) (idCustomer int, usernameCustomer string)
	Logout(ctx context.Context, customerId int)
	Update(ctx context.Context, request web.CustomerUpdateRequest) web.CustomerResponse
	Delete(ctx context.Context, customerId int)
	FindById(ctx context.Context, customerId int) web.CustomerResponse
	FindAll(ctx context.Context) []web.CustomerResponse
}
