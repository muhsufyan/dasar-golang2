package service

import (
	"context"

	"github.com/muhsufyan/dasar-golang2/model/web"
)

type CategoryService interface {
	/*
		ketika Create maka requestnya apa ? karena tdk mungkin domain (domain akan terhubung dg repository bukan service). sehingga kita buat model baru yg akan menjd request dr service
		dan untuk response nya kita perlu buat juga data model response (sama sprti request)
	*/
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}
