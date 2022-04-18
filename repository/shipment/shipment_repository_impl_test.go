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
	shipment := entity.Shipment{
		Adress: "jl tanpa bayangan",
		Cost:   "100000",
	}

	result, err := shipmentRepository.Insert(ctx, shipment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	shipmentRepository := NewShipmentRepository(belajar_golang.GetConnection())

	shipment, err := shipmentRepository.FindById(context.Background(), 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(shipment)
}

func TestFindAll(t *testing.T) {
	shipmentRepository := NewShipmentRepository(belajar_golang.GetConnection())

	shipments, err := shipmentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, shipment := range shipments {
		fmt.Println(shipment)
	}
}

func TestUpdate(t *testing.T) {
	shipmentRepository := NewShipmentRepository(belajar_golang.GetConnection())

	ctx := context.Background()
	shipment := entity.Shipment{
		Adress: "jl tanpa bayangan",
		Cost:   "900000",
	}

	result, err := shipmentRepository.Update(ctx, shipment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestDelete(t *testing.T) {
	shipmentRepository := NewShipmentRepository(belajar_golang.GetConnection())

	ctx := context.Background()
	shipment := entity.Shipment{
		Adress: "jl tanpa bayangan",
	}

	result, err := shipmentRepository.Delete(ctx, shipment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
