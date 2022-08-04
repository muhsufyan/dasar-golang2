package service

import (
	"context"
	"database/sql"

	"github.com/muhsufyan/dasar-golang2/helper"
	"github.com/muhsufyan/dasar-golang2/model/domain"
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
	defer helper.CommitOrRollback(tx)

	// implementasi busines logic create category
	// get name
	category := domain.Category{
		Name: request.Name,
	}
	// jalankan query lewat repository
	category = service.CategoryRepository.Save(ctx, tx, category)
	// agar data yg dikembalikan ke user/client berupa data model response maka kita perlu buat dan panggil fungsi yg ktia buat
	// untuk convert dari data model input ke db (lewat domain) menjd data model response (lewat helper)
	return helper.Convert2CategoryResponse(category)
}
func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	// mulai transaksi ke db
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	// implementasi busines logic create category
	// 1. cek id category yg akan di update
	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)
	// simpan nama inputan dari request user
	category.Name = request.Name
	//2. jalankan query update lewat repository
	category = service.CategoryRepository.Update(ctx, tx, category)
	return helper.Convert2CategoryResponse(category)
}
func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	// mulai transaksi ke db
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// implementasi busines logic create category
	// cek id category yg akan di update
	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)
	// jalankan query delete lewat repository
	service.CategoryRepository.Delete(ctx, tx, category)
}
func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	// mulai transaksi ke db
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	// implementasi busines logic create category
	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)
	return helper.Convert2CategoryResponse(category)
}
func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	// mulai transaksi ke db
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	// implementasi busines logic create category
	categories := service.CategoryRepository.FindAll(ctx, tx)
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, helper.Convert2CategoryResponse(category))
	}
	return categoryResponses
}