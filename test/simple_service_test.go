package test

import (
	"fmt"
	"testing"

	"github.com/muhsufyan/dasar-golang2/simple"
	"github.com/stretchr/testify/assert"
)

func TestSimpleService(t *testing.T) {
	// buat test simple service hasil dari wire injector tinggal panggil func InitializedService
	// tambah error
	simpleService, err := simple.InitializedService(false)
	fmt.Println(simpleService.SimpleRepository)
	fmt.Println(err)
	assert.NotNil(t, simpleService)
	assert.Nil(t, err)
}

func TestSimpleServiceFailedError(t *testing.T) {
	// buat test simple service hasil dari wire injector tinggal panggil func InitializedService
	// tambah error
	simpleService, err := simple.InitializedService(true)
	fmt.Println(simpleService)
	fmt.Println(err)
	assert.Nil(t, simpleService)
	assert.NotNil(t, err)
}
