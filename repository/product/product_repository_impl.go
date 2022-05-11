package repository

// implenteasi adalah bagaimana method(dari repositori) itu dikeksekusi

import (
	"context"
	"database/sql"
	"errors"
	"go-database/entity"
	"strconv"
)

type productRepositoryImpl struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepositoryImpl{DB: db}
}

func (repository *productRepositoryImpl) Insert(ctx context.Context, name entity.Product) (entity.Product, error) {
	script := "INSERT INTO product(name,harga) VALUES (?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, name.Name, name.Harga)
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

func (repository *productRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Product, error) {
	script := "SELECT id, name, harga FROM product WHERE id = ? Limit 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	product := entity.Product{}
	if err != nil {
		return product, err
	}
	defer rows.Close()
	if rows.Next() {
		// ada
		rows.Scan(&product.Id, &product.Name, &product.Harga)
		return product, nil
	} else {
		// tidak ada
		return product, errors.New("Id " + strconv.Itoa(int(id)) + " Not Found")
	}
}
func (repository *productRepositoryImpl) FindAll(ctx context.Context) ([]entity.Product, error) {
	script := "SELECT id, name, harga FROM product"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []entity.Product
	for rows.Next() {
		product := entity.Product{}
		rows.Scan(&product.Id, &product.Name, &product.Harga)
		products = append(products, product)
	}
	return products, nil
}

func (repository *productRepositoryImpl) Update(ctx context.Context, product entity.Product) (entity.Product, error) {
	script := "UPDATE product SET phone = ?, name = ? WHERE id = ?"
	result, err := repository.DB.ExecContext(ctx, script, product.Name, product.Harga)
	if err != nil {
		return product, err
	}
	rowCnt, err := result.RowsAffected()
	if err != nil {
		return product, err
	}
	if rowCnt == 0 {
		return product, err
	}
	return product, err
}

func (repository *productRepositoryImpl) Delete(ctx context.Context, id int32) (int32, error) {
	script := "DELETE FROM product WHERE id = ?"
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
