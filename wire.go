//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/Genarodaniel/go-DI/product"
	"github.com/google/wire"
)

var setRepositoryDependency = wire.NewSet(
	product.NewProductRepository,
	wire.Bind(new(product.IProductRepository), new(*product.ProductRepository)),
)

func NewProductUseCase(db *sql.DB) *product.ProductUseCase {
	wire.Build(
		setRepositoryDependency,
		product.NewProductUseCase,
	)

	return &product.ProductUseCase{}
}
