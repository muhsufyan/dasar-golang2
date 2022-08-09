package simple

// interface
type SayHallo interface {
	// kontrak
	Hallo(name string) string
}
type HalloService struct {
	SayHallo
}

// implementasi
type SayHalloImp struct {
}

// func implementasi
func (s *SayHalloImp) Hallo(name string) string {
	return "Hallo" + name
}

// provider. return implementasi
func NewSayHalloImp() *SayHalloImp {
	return &SayHalloImp{}
}

// provider. return interface & parameternya berupa interface
func NewHalloService(sayHello SayHallo) *HalloService {
	return &HalloService{SayHallo: sayHello}
}
