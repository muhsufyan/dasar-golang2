package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/muhsufyan/dasar-golang2/helper"
	"github.com/muhsufyan/dasar-golang2/model/domain"
)

//implementasi dari contruct interface category repository
type CategoryRepositoryImpl struct {
}

// contructor untuk repository, kita tdk perlu param apapun (sesuai dg struct CategoryRepositoryImpl yg tdk menangkap/menyimpan data)
// return-nya interface repository tp sama sprti constructor untuk repository yg dikembalikan sbnrnya adlh struct dari implementasi-nya (dibhs lain ini disbt dg polymorphisme)
// func NewCategoryRepository() CategoryRepository {
// 	return &CategoryRepositoryImpl{}
// }

// provider diatas ubah returnnya
func NewCategoryRepository() *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{}
}
func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	sql := "INSERT INTO category(name) VALUE (?)"
	result, err := tx.ExecContext(ctx, sql, category.Name)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	category.Id = int(id)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	sql := "UPDATE category SET name = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, sql, category.Name, category.Id)
	helper.PanicIfError(err)
	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	sql := "DELETE FROM category WHERE id = ?"
	_, err := tx.ExecContext(ctx, sql, category.Id)
	helper.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	sql := "SELECT id, name FROM category WHERE id = ?"
	rows, err := tx.QueryContext(ctx, sql, categoryId)
	helper.PanicIfError(err)
	defer rows.Close()
	// untuk menampung data
	category := domain.Category{}
	// jika datanya ada
	if rows.Next() {
		// catch data result query and store to category
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("category not found")
	}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	sql := "SELECT id, name FROM category"
	rows, err := tx.QueryContext(ctx, sql)
	helper.PanicIfError(err)
	defer rows.Close()
	// untuk menyimpan data
	var categories []domain.Category
	// mapping data
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}
