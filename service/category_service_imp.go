package service

import (
	"context"
	"database/sql"

	"github.com/muhsufyan/dasar-golang2/helper"
	"github.com/muhsufyan/dasar-golang2/model/web"
	"github.com/muhsufyan/dasar-golang2/repository"
)

type CategoryServiceImpl struct {
	// injek dg interface dr CategoryRepository
	CategoryRepository repository.CategoryRepository
	// injek dg db
	DB *sql.DB
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	// mulai transaksi untuk create category
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer func() {
		err := recover()
		if err != nil {
			tx.Rollback()
			panic(err)
		} else {
			tx.Commit()
		}
	}()
}
