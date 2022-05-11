package repository

// implenteasi adalah bagaimana method(dari repositori) itu dikeksekusi

import (
	"context"
	"database/sql"
	"errors"
	"go-database/entity"
	"strconv"
)

type paymentRepositoryImpl struct {
	DB *sql.DB
}

func NewPaymentRepository(db *sql.DB) PaymentRepository {
	return &paymentRepositoryImpl{DB: db}
}

func (repository *paymentRepositoryImpl) Insert(ctx context.Context, method entity.Payment) (entity.Payment, error) {
	script := "INSERT INTO payment(method,bank_name) VALUES (?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, method.Method, method.Bank)
	if err != nil {
		return method, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return method, err
	}
	method.Id = int32(id)
	return method, nil
}

func (repository *paymentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Payment, error) {
	script := "SELECT id, method, bank_name FROM payment WHERE id = ? Limit 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	payment := entity.Payment{}
	if err != nil {
		return payment, err
	}
	defer rows.Close()
	if rows.Next() {
		// ada
		rows.Scan(&payment.Id, &payment.Method, &payment.Bank)
		return payment, nil
	} else {
		// tidak ada
		return payment, errors.New("Id " + strconv.Itoa(int(id)) + " Not Found")
	}
}
func (repository *paymentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Payment, error) {
	script := "SELECT id, method, bank_name FROM payment"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var payments []entity.Payment
	for rows.Next() {
		payment := entity.Payment{}
		rows.Scan(&payment.Id, &payment.Method, &payment.Bank)
		payments = append(payments, payment)
	}
	return payments, nil
}

func (repository *paymentRepositoryImpl) Update(ctx context.Context, payment entity.Payment) (entity.Payment, error) {
	script := "UPDATE payment SET bank_name = ?, method = ? WHERE id = ?"
	result, err := repository.DB.ExecContext(ctx, script, payment.Bank, payment.Method, payment.Id)
	if err != nil {
		return payment, err
	}
	rowCnt, err := result.RowsAffected()
	if err != nil {
		return payment, err
	}
	if rowCnt == 0 {
		return payment, err
	}
	return payment, err
}

func (repository *paymentRepositoryImpl) Delete(ctx context.Context, id int32) (int32, error) {
	script := "DELETE FROM payment WHERE id = ?"
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
