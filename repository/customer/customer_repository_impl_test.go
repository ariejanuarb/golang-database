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
	customerRepository := NewCustomerRepository(belajar_golang.GetConnection())

	ctx := context.Background()
	customer := entity.Customer{
		Name:  "Sonia",
		Phone: "0816262626",
	}

	result, err := customerRepository.Insert(ctx, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	customerRepository := NewCustomerRepository(belajar_golang.GetConnection())

	customer, err := customerRepository.FindById(context.Background(), 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
}

func TestFindAll(t *testing.T) {
	customerRepository := NewCustomerRepository(belajar_golang.GetConnection())

	customers, err := customerRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, customer := range customers {
		fmt.Println(customer)
	}
}

func TestUpdate(t *testing.T) {
	customerRepository := NewCustomerRepository(belajar_golang.GetConnection())

	ctx := context.Background()
	customer := entity.Customer{
		Name:  "Arie",
		Phone: "0812345678",
	}

	result, err := customerRepository.Update(ctx, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestDelete(t *testing.T) {
	customerRepository := NewCustomerRepository(belajar_golang.GetConnection())

	ctx := context.Background()
	customer := entity.Customer{
		Name: "Sonia",
	}

	result, err := customerRepository.Delete(ctx, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
