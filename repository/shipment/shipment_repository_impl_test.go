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
	shipmentRepository := NewShipmentRepository(belajar_golang.GetConnection())

	ctx := context.Background()
	customer := entity.Shipment{
		Adress: "jl tanpa bayangan",
		Cost:   "100000",
	}

	result, err := shipmentRepository.Insert(ctx, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	shipmentRepository := NewShipmentRepository(belajar_golang.GetConnection())

	comment, err := shipmentRepository.FindById(context.Background(), 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	shipmentRepository := NewShipmentRepository(belajar_golang.GetConnection())

	comments, err := shipmentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}

func TestUpdate(t *testing.T) {
	shipmentRepository := NewShipmentRepository(belajar_golang.GetConnection())

	ctx := context.Background()
	customer := entity.Shipment{
		Adress: "jl tanpa bayangan",
		Cost:   "900000",
	}

	result, err := shipmentRepository.Update(ctx, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestDelete(t *testing.T) {
	shipmentRepository := NewShipmentRepository(belajar_golang.GetConnection())

	ctx := context.Background()
	customer := entity.Shipment{
		Adress: "jl tanpa bayangan",
	}

	result, err := shipmentRepository.Delete(ctx, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
