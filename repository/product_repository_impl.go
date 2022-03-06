package repository

import (
	"belajar-golang-rest-api/helper"
	"belajar-golang-rest-api/model/domain"
	"context"
	"database/sql"
	"errors"
)

type ProductRepositoryImpl struct {
}

func (repository *ProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "INSERT INTO product(name) values (?)"
	result, err := tx.ExecContext(ctx, SQL, product.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	product.Id = int(id)
	return product
}

func (repository *ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "UPDATE product SET name = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, product.Name, product.Id)
	helper.PanicIfError(err)

	return product
}

func (repository *ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, productId int) {
	SQL := "DELETE FROM product where id = ?"
	_, err := tx.ExecContext(ctx, SQL, productId)
	helper.PanicIfError(err)
}

func (repository *ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId int) (domain.Product, error) {
	SQL := "SELECT id, name FROM product WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, productId)
	helper.PanicIfError(err)
	defer rows.Close()

	product := domain.Product{}

	if rows.Next() {
		err := rows.Scan(&product.Id, &product.Name)
		helper.PanicIfError(err)
		return product, nil
	} else {
		return product, errors.New("product is not found")
	}
}

func (repository *ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Product {
	SQL := "SELECT id, name FROM product"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		product := &domain.Product{}
		err := rows.Scan(&product.Id, &product.Name)
		helper.PanicIfError(err)
		products = append(products, *product)
	}

	return products
}
