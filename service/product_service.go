package service

import (
	"belajar-golang-rest-api/model/web"
	"context"
)

type ProductService interface {
	Create(ctx context.Context, request web.ProductCreateRequest) web.ProductResponse
	Update(ctx context.Context, request web.ProductUpdateRequest) web.ProductResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) web.ProductResponse
	FindAll(ctx context.Context) []web.ProductResponse
}
