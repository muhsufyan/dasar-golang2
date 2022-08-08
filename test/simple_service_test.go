package test

import (
	"fmt"
	"testing"

	"github.com/muhsufyan/dasar-golang2/simple"
)

func TestSimpleService(t *testing.T) {
	// buat test simple service hasil dari wire injector tinggal panggil func InitializedService
	simpleService := simple.InitializedService()
	fmt.Println(simpleService.SimpleRepository)
}
