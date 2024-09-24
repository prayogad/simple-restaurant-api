package repository

import (
	"context"
	"database/sql"
	"errors"
	"simple-restaurant-web/helper"
	"simple-restaurant-web/model/domain"
	"strconv"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type CustomerRepositoryImpl struct {
}

func NewCustomerRepository() CustomerRepository {
	return &CustomerRepositoryImpl{}
}

func (repository *CustomerRepositoryImpl) Save(ctx context.Context, Tx *sql.Tx, customer domain.Customer) (domain.Customer, error) {
	Query := "SELECT id FROM customer WHERE username = $1"
	rows, err := Tx.QueryContext(ctx, Query, customer.Username)
	helper.PanicIfError(err)
	defer rows.Close()

	if rows.Next() {
		return customer, errors.New("username already taken")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(customer.Password), bcrypt.DefaultCost)
	helper.PanicIfError(err)
	customer.Password = string(hashedPassword)

	SQL := "INSERT INTO customer(username, password) VALUES ($1, $2) RETURNING id"

	var id int
	Tx.QueryRowContext(ctx, SQL, customer.Username, customer.Password).Scan(&id)
	customer.Id = id

	return customer, nil
}

func (repository *CustomerRepositoryImpl) Login(ctx context.Context, Tx *sql.Tx, username string, password string) (domain.Customer, error) {
	SQL := "SELECT id, username, password FROM customer WHERE username = $1"

	rows, err := Tx.QueryContext(ctx, SQL, username)
	helper.PanicIfError(err)
	defer rows.Close()

	customer := domain.Customer{}
	if rows.Next() {
		err := rows.Scan(&customer.Id, &customer.Username, &customer.Password)
		helper.PanicIfError(err)
	} else {
		return customer, errors.New("customer not found")
	}

	passwordErr := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(password))
	if passwordErr != nil {
		return customer, errors.New("username or password incorrect")
	}

	token := uuid.New().String()
	customer.Token = token

	SQL = "UPDATE customer SET token = '" + customer.Token + "' WHERE username = '" + customer.Username + "' "
	Tx.ExecContext(ctx, SQL)

	return customer, nil
}

func (repository *CustomerRepositoryImpl) ValidateToken(Tx *sql.Tx, token string) (idCustomer int, usernameCustomer string) {
	SQL := "SELECT id, username FROM customer WHERE token = $1"
	rows, err := Tx.Query(SQL, token)
	helper.PanicIfError(err)
	defer rows.Close()

	var username, id string

	if rows.Next() {
		err := rows.Scan(&id, &username)
		helper.PanicIfError(err)
	} else {
		return 0, ""
	}
	newId, _ := strconv.Atoi(id)
	return newId, username
}

func (repository *CustomerRepositoryImpl) Logout(ctx context.Context, Tx *sql.Tx, customerId int) {
	SQL := "UPDATE customer SET token = null WHERE id = $1"
	_, err := Tx.ExecContext(ctx, SQL, customerId)
	helper.PanicIfError(err)

}

func (repository *CustomerRepositoryImpl) Update(ctx context.Context, Tx *sql.Tx, customer domain.Customer) domain.Customer {
	currentCustomer := domain.Customer{}
	currentCustomer.Id = customer.Id
	Query := "SELECT username, password FROM customer WHERE id = $1"
	err := Tx.QueryRowContext(ctx, Query, currentCustomer.Id).Scan(&currentCustomer.Username, &currentCustomer.Password)
	helper.PanicIfError(err)

	if customer.Username != "" {
		currentCustomer.Username = customer.Username
	}

	if customer.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(customer.Password), bcrypt.DefaultCost)
		helper.PanicIfError(err)
		currentCustomer.Password = string(hashedPassword)
	}

	SQL := "UPDATE customer SET username = $1, password = $2 WHERE id = $3"
	_, err = Tx.ExecContext(ctx, SQL, currentCustomer.Username, currentCustomer.Password, customer.Id)
	helper.PanicIfError(err)

	return currentCustomer
}

func (repository *CustomerRepositoryImpl) Delete(ctx context.Context, Tx *sql.Tx, customer domain.Customer) {
	SQL := "DELETE FROM customer WHERE id = $1"
	_, err := Tx.ExecContext(ctx, SQL, customer.Id)
	helper.PanicIfError(err)
}

func (repository *CustomerRepositoryImpl) FindById(ctx context.Context, Tx *sql.Tx, customerId int) (domain.Customer, error) {
	SQL := "SELECT id, username FROM customer WHERE id = $1"
	rows, err := Tx.QueryContext(ctx, SQL, customerId)
	helper.PanicIfError(err)
	defer rows.Close()

	customer := domain.Customer{}
	if rows.Next() {
		err := rows.Scan(&customer.Id, &customer.Username)
		helper.PanicIfError(err)
		return customer, nil
	} else {
		return customer, errors.New("customer not found")
	}
}

func (repository *CustomerRepositoryImpl) FindAll(ctx context.Context, Tx *sql.Tx) []domain.Customer {
	SQL := "SELECT id, username FROM customer"
	rows, err := Tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var customers []domain.Customer
	for rows.Next() {
		customer := domain.Customer{}
		err := rows.Scan(&customer.Id, &customer.Username)
		helper.PanicIfError(err)
		customers = append(customers, customer)
	}

	return customers
}
