package repository

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	belajar_golang "go-database"
	"go-database/entity"
	"testing"
)

func TestInsert(t *testing.T) {
	invoiceRepository := NewInvoiceRepository(belajar_golang.GetConnection())

	ctx := context.Background()
	invoice := entity.Invoice{
		Number: "123456",
		Date:   "2022-01-01",
	}

	result, err := invoiceRepository.Insert(ctx, invoice)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	invoiceRepository := NewInvoiceRepository(belajar_golang.GetConnection())

	invoice, err := invoiceRepository.FindById(context.Background(), 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(invoice)
}

func TestFindAll(t *testing.T) {
	invoiceRepository := NewInvoiceRepository(belajar_golang.GetConnection())

	invoices, err := invoiceRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, invoice := range invoices {
		fmt.Println(invoice)
	}
}

func TestUpdate(t *testing.T) {
	invoiceRepository := NewInvoiceRepository(belajar_golang.GetConnection())

	ctx := context.Background()
	invoice := entity.Invoice{
		Date:   "2022-01-03",
		Number: "123456",
	}

	result, err := invoiceRepository.Update(ctx, invoice)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestDelete(t *testing.T) {
	invoiceRepository := NewInvoiceRepository(belajar_golang.GetConnection())

	ctx := context.Background()
	invoice := entity.Invoice{
		Number: "123456",
	}

	result, err := invoiceRepository.Delete(ctx, invoice)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
