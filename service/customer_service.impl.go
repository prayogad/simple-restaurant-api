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

type CustomerServiceImpl struct {
	CustomerRepository repository.CustomerRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCustomerService(customerRepository repository.CustomerRepository, DB *sql.DB, validate *validator.Validate) CustomerService {
	return &CustomerServiceImpl{
		CustomerRepository: customerRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *CustomerServiceImpl) Create(ctx context.Context, request web.CustomerCreateRequest) web.CustomerResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customer := domain.Customer{
		Username: request.Username,
		Password: request.Password,
	}

	customer, err = service.CustomerRepository.Save(ctx, tx, customer)
	if err != nil {
		panic(exceptions.NewUsernameTakenError(err.Error()))
	}

	return helper.ToCustomerResponse(customer)
}

func (service *CustomerServiceImpl) Login(ctx context.Context, request web.CustomerLoginRequest) web.CustomerResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customer, err := service.CustomerRepository.Login(ctx, tx, request.Username, request.Password)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	return helper.ToCustomerLoginResponse(customer)
}

func (service *CustomerServiceImpl) ValidateToken(token string) (idCustomer int, usernameCustomer string) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	id, username := service.CustomerRepository.ValidateToken(tx, token)

	return id, username
}

func (service *CustomerServiceImpl) Logout(ctx context.Context) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_, err = service.CustomerRepository.CurrentCustomer(ctx, tx)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	service.CustomerRepository.Logout(ctx, tx)
}

func (service *CustomerServiceImpl) Update(ctx context.Context, request web.CustomerUpdateRequest) web.CustomerResponse {
	request.Id = ctx.Value("idCustomer").(int)
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customer, err := service.CustomerRepository.CurrentCustomer(ctx, tx)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	customer.Username = request.Username
	customer.Password = request.Password

	customer = service.CustomerRepository.Update(ctx, tx, customer)

	return helper.ToCustomerResponse(customer)
}

func (service *CustomerServiceImpl) Delete(ctx context.Context, customerId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customer, err := service.CustomerRepository.CurrentCustomer(ctx, tx)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	service.CustomerRepository.Delete(ctx, tx, customer)
}

func (service *CustomerServiceImpl) CurrentCustomer(ctx context.Context) web.CustomerResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customer, err := service.CustomerRepository.CurrentCustomer(ctx, tx)
	if err != nil {
		panic(exceptions.NewNotFoundError(err.Error()))
	}

	return helper.ToCustomerResponse(customer)
}

func (service *CustomerServiceImpl) FindAll(ctx context.Context) []web.CustomerResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customers := service.CustomerRepository.FindAll(ctx, tx)

	return helper.ToCustomerResponses(customers)
}
