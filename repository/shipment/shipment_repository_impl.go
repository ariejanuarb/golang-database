package repository

// implenteasi adalah bagaimana method(dari repositori) itu dikeksekusi

import (
	"context"
	"database/sql"
	"errors"
	"go-database/entity"
	"strconv"
)

type shipmentRepositoryImpl struct {
	DB *sql.DB
}

func NewShipmentRepository(db *sql.DB) ShipmentRepository {
	return &shipmentRepositoryImpl{DB: db}
}

func (repository *shipmentRepositoryImpl) Insert(ctx context.Context, adress entity.Shipment) (entity.Shipment, error) {
	script := "INSERT INTO shipment(adress,cost) VALUES (?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, adress.Adress, adress.Cost)
	if err != nil {
		return adress, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return adress, err
	}
	adress.Id = int32(id)
	return adress, nil
}

func (repository *shipmentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Shipment, error) {
	script := "SELECT id, adress, cost FROM shipment WHERE id = ? Limit 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	shipment := entity.Shipment{}
	if err != nil {
		return shipment, err
	}
	defer rows.Close()
	if rows.Next() {
		// ada
		rows.Scan(&shipment.Id, &shipment.Adress, &shipment.Cost)
		return shipment, nil
	} else {
		// tidak ada
		return shipment, errors.New("Id " + strconv.Itoa(int(id)) + " Not Found")
	}
}
func (repository *shipmentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Shipment, error) {
	script := "SELECT id, adress, cost FROM shipment"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var shipments []entity.Shipment
	for rows.Next() {
		shipment := entity.Shipment{}
		rows.Scan(&shipment.Id, &shipment.Adress, &shipment.Cost)
		shipments = append(shipments, shipment)
	}
	return shipments, nil
}

func (repository *shipmentRepositoryImpl) Update(ctx context.Context, shipment entity.Shipment) (entity.Shipment, error) {
	script := "UPDATE shipment SET cost = ?, adress = ? WHERE id = ?"
	result, err := repository.DB.ExecContext(ctx, script, shipment.Adress, shipment.Cost)
	if err != nil {
		return shipment, err
	}
	rowCnt, err := result.RowsAffected()
	if err != nil {
		return shipment, err
	}
	if rowCnt == 0 {
		return shipment, err
	}
	return shipment, err
}

func (repository *shipmentRepositoryImpl) Delete(ctx context.Context, id int32) (int32, error) {
	script := "DELETE FROM shipment WHERE id = ?"
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
