package repository

import (
	"context"
	"database/sql"

	"github.com/muhsufyan/dasar-golang2/model/domain"
)

// buat contruct dlm bntk interface, agar tahu detail contruct (dlm hal ini category)
type CategoryRepository interface {
	// operasi/method yg dpt dilakukan. method ini nantinya berupa query. param hrs diawali dg context. kita gunakan db transaksional meskipun ini seharusnya bukan transaksi
	// param ke3 adlh data aslinya
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
}
