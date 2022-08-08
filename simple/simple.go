package simple

import "errors"

// buat objek repository
type SimpleRepository struct {
	// tambah error untuk handle error provider
	Error bool
}

// buat provider repository (buat func provicer repository). ralat func ini adalah injector
// add injector parameter yaitu isError
func NewSimpleRepository(isError bool) *SimpleRepository {
	return &SimpleRepository{
		//injector parameter
		Error: isError,
	}
}

// buat objek service
// service is depend ke respository (service perlu diinjek oleh repository)
type SimpleService struct {
	// service diinjek dg struct repository
	*SimpleRepository
}

// buat provider service yg perlu (depend) repository, jd repository dijdkan as param karena service depend ke repository. (buat func provicer servie)
// untuk handle error provider kita tambahkan return valunya error
func NewSimpleService(repository *SimpleRepository) (*SimpleService, error) {
	// tambah if kondisi untuk handle provider error pd constructor provider repository
	// jika ada error
	if repository.Error {
		return nil, errors.New("failed create service")
	} else {
		return &SimpleService{
			SimpleRepository: repository,
		}, nil
	}
}
