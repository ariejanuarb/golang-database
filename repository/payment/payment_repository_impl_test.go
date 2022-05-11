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
	paymentRepository := NewPaymentRepository(belajar_golang.GetConnection())

	ctx := context.Background()
	payment := entity.Payment{
		Method: "teller",
		Bank:   "Bank Syariah Indonesia",
	}

	result, err := paymentRepository.Insert(ctx, payment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	paymentRepository := NewPaymentRepository(belajar_golang.GetConnection())

	payment, err := paymentRepository.FindById(context.Background(), 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(payment)
}

func TestFindAll(t *testing.T) {
	paymentRepository := NewPaymentRepository(belajar_golang.GetConnection())

	payments, err := paymentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, payment := range payments {
		fmt.Println(payment)
	}
}

func TestUpdate(t *testing.T) {
	paymentRepository := NewPaymentRepository(belajar_golang.GetConnection())

	ctx := context.Background()
	payment := entity.Payment{
		Id:     1,
		Method: "COD baru",
		Bank:   "gopay baru",
	}

	result, err := paymentRepository.Update(ctx, payment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestDelete(t *testing.T) {
	paymentRepository := NewPaymentRepository(belajar_golang.GetConnection())

	result, err := paymentRepository.Delete(context.Background(), 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
