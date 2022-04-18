package repository

// implenteasi adalah bagaimana method(dari repositori) itu dikeksekusi

import (
	"context"
	"database/sql"
	"errors"
	"go-database/entity"
	"strconv"
)

type invoiceRepositoryImpl struct {
	DB *sql.DB
}

func NewInvoiceRepository(db *sql.DB) InvoiceRepository {
	return &invoiceRepositoryImpl{DB: db}
}

func (repository *invoiceRepositoryImpl) Insert(ctx context.Context, no_invoice entity.Invoice) (entity.Invoice, error) {
	script := "INSERT INTO invoice(no_invoice,invoice_date) VALUES (?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, no_invoice.Number, no_invoice.Date)
	if err != nil {
		return no_invoice, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return no_invoice, err
	}
	no_invoice.Id = int32(id)
	return no_invoice, nil
}

// service = test, repo = impl
func (repository *invoiceRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Invoice, error) {
	script := "SELECT id, no_invoice, invoice_date FROM invoice WHERE id = ? Limit 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	invoice := entity.Invoice{}
	if err != nil {
		return invoice, err
	}
	defer rows.Close()
	if rows.Next() {
		// ada
		rows.Scan(&invoice.Id, &invoice.Number, &invoice.Date)
		return invoice, nil
	} else {
		// tidak ada
		return invoice, errors.New("Id " + strconv.Itoa(int(id)) + " Not Found")
	}
}
func (repository *invoiceRepositoryImpl) FindAll(ctx context.Context) ([]entity.Invoice, error) {
	script := "SELECT id, no_invoice, invoice_date FROM invoice"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var invoices []entity.Invoice
	for rows.Next() {
		invoice := entity.Invoice{}
		rows.Scan(&invoice.Id, &invoice.Number, &invoice.Date)
		invoices = append(invoices, invoice)
	}
	return invoices, nil
}

func (repository *invoiceRepositoryImpl) Update(ctx context.Context, invoice entity.Invoice) (entity.Invoice, error) {
	script := "UPDATE invoice SET invoice_date = ? WHERE no_invoice = ?"
	result, err := repository.DB.ExecContext(ctx, script, invoice.Date, invoice.Number)
	if err != nil {
		return invoice, err
	}
	rowCnt, err := result.RowsAffected()
	if err != nil {
		return invoice, err
	}
	if rowCnt == 0 {
		return invoice, err
	}
	return invoice, err
}

func (repository *invoiceRepositoryImpl) Delete(ctx context.Context, invoice entity.Invoice) (entity.Invoice, error) {
	script := "DELETE FROM invoice WHERE no_invoice = ?"
	result, err := repository.DB.ExecContext(ctx, script, invoice.Number)
	if err != nil {
		return invoice, err
	}
	rowCnt, err := result.RowsAffected()
	if err != nil {
		return invoice, err
	}
	if rowCnt == 0 {
		return invoice, err
	}
	return invoice, nil
}
