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
	productRepository := NewProductRepository(belajar_golang.GetConnection())

	ctx := context.Background()
	product := entity.Product{
		Name:  "Meja Coding",
		Harga: "600000",
	}

	result, err := productRepository.Insert(ctx, product)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	productRepository := NewProductRepository(belajar_golang.GetConnection())

	product, err := productRepository.FindById(context.Background(), 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(product)
}

func TestFindAll(t *testing.T) {
	productRepository := NewProductRepository(belajar_golang.GetConnection())

	products, err := productRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, product := range products {
		fmt.Println(product)
	}
}

func TestUpdate(t *testing.T) {
	productRepository := NewProductRepository(belajar_golang.GetConnection())

	ctx := context.Background()
	product := entity.Product{
		Id:    1,
		Name:  "sling bag baru",
		Harga: "900000",
	}

	result, err := productRepository.Update(ctx, product)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestDelete(t *testing.T) {
	productRepository := NewProductRepository(belajar_golang.GetConnection())

	result, err := productRepository.Delete(context.Background(), 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
