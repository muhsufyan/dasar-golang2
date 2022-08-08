package test

import (
	"fmt"
	"testing"

	"github.com/muhsufyan/dasar-golang2/simple"
)

// set simple/simple.go dimana Provider repository (NewSimpleRepository) return Errornya adlh false/bisa juga ckp return &SimpleRepository{}
func TestSimpleService(t *testing.T) {
	// buat test simple service hasil dari wire injector tinggal panggil func InitializedService
	// tambah error
	simpleService, err := simple.InitializedService()
	fmt.Println(simpleService.SimpleRepository)
	fmt.Println(err)
}

// set simple/simple.go dimana Provider repository (NewSimpleRepository) return Errornya adlh true
func TestSimpleServiceFailedError(t *testing.T) {
	// buat test simple service hasil dari wire injector tinggal panggil func InitializedService
	// tambah error
	simpleService, err := simple.InitializedService()
	fmt.Println(simpleService)
	fmt.Println(err)
}
