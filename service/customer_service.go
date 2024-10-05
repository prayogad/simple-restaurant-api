package service

import (
	"context"
	"simple-restaurant-web/model/web"
)

type CustomerService interface {
	Create(ctx context.Context, request web.CustomerCreateRequest) web.CustomerResponse
	Login(ctx context.Context, request web.CustomerLoginRequest) web.CustomerResponse
	ValidateToken(token string) (idCustomer int, usernameCustomer string)
	Logout(ctx context.Context)
	Update(ctx context.Context, request web.CustomerUpdateRequest) web.CustomerResponse
	Delete(ctx context.Context, customerId int)
	CurrentCustomer(ctx context.Context) web.CustomerResponse
	FindAll(ctx context.Context) []web.CustomerResponse
}
