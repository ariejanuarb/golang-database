package repository

// implenteasi adalah bagaimana method(dari repositori) itu dikeksekusi

import (
	"context"
	"database/sql"
	"errors"
	"go-database/entity"
	"strconv"
)

type customerRepositoryImpl struct {
	DB *sql.DB
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &customerRepositoryImpl{DB: db}
}

func (repository *customerRepositoryImpl) Insert(ctx context.Context, name entity.Customer) (entity.Customer, error) {
	script := "INSERT INTO customer(name,phone) VALUES (?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, name.Name, name.Phone)
	if err != nil {
		return name, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return name, err
	}
	name.Id = int32(id)
	return name, nil
}

func (repository *customerRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Customer, error) {
	script := "SELECT id, name, phone FROM customer WHERE id = ? Limit 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	customer := entity.Customer{}
	if err != nil {
		return customer, err
	}
	defer rows.Close()
	if rows.Next() {
		// ada
		rows.Scan(&customer.Id, &customer.Name, &customer.Phone)
		return customer, nil
	} else {
		// tidak ada
		return customer, errors.New("Id " + strconv.Itoa(int(id)) + " Not Found")
	}
}
func (repository *customerRepositoryImpl) FindAll(ctx context.Context) ([]entity.Customer, error) {
	script := "SELECT id, name, phone FROM customer"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var customers []entity.Customer
	for rows.Next() {
		customer := entity.Customer{}
		rows.Scan(&customer.Id, &customer.Name, &customer.Phone)
		customers = append(customers, customer)
	}
	return customers, nil
}

func (repository *customerRepositoryImpl) Update(ctx context.Context, customer entity.Customer) (entity.Customer, error) {
	script := "UPDATE customer SET phone = ?, name = ? WHERE id = ?"
	result, err := repository.DB.ExecContext(ctx, script, customer.Phone, customer.Name, customer.Id)
	if err != nil {
		return customer, err
	}
	rowCnt, err := result.RowsAffected()
	if err != nil {
		return customer, err
	}
	if rowCnt == 0 {
		return customer, err
	}
	return customer, err
}

func (repository *customerRepositoryImpl) Delete(ctx context.Context, id int32) (int32, error) {
	script := "DELETE FROM customer WHERE id = ?"
	result, err := repository.DB.ExecContext(ctx, script, id)
	if err != nil {
		return id, err
	}
	rowCnt, err := result.RowsAffected()
	if err != nil {
		return id, err
	}
	if rowCnt == 0 {
		return id, err
	}
	return id, nil
}
