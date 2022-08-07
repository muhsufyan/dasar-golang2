package simple

// buat objek repository
type SimpleRepository struct {
}

// buat provider repository (buat func provicer repository)
func NewSimpleRepository() *SimpleRepository {
	return &SimpleRepository{}
}

// buat objek service
// service is depend ke respository (service perlu diinjek oleh repository)
type SimpleService struct {
	// service diinjek dg struct repository
	*SimpleRepository
}

// buat provider service yg perlu (depend) repository, jd repository dijdkan as param karena service depend ke repository. (buat func provicer servie)
func NewSimpleService(repository *SimpleRepository) *SimpleService {
	return &SimpleService{SimpleRepository: repository}
}
