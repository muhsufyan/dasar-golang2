//go:build wireinject
// +build wireinject

package simple

import "github.com/google/wire"

// buat func injector
// return-nya adlh dependency terakhir yg diperlukan (dlm kasus ini adlh service), nama func injector kita awali dg Initialize....
// ini mirip sprti Constructor (New....)
// agar injector ini dpt handle error (dampaknya/sehingga provider dpt handle error) maka tambah return value berupa error
// jdkan isError as parameter provider
func InitializedService(isError bool) (*SimpleService, error) {
	// masukkan provider apa saja yg akan diinjek/digunakan
	// INGAT FUNC MANA YG PERLU DIINJEK AKAN DIHANDLE OLEH GOOGLE WIRE, PD KASUS INI service perlu diinjek oleh repository NAH HAL INI AKAN DIHANDLE OLEH WIRE
	// JD KITA HANYA MENDAFTARKAN FUNC PROVIDER YG AKAN DIINJEK/DIGUNAKAN SAJA
	wire.Build(
		NewSimpleRepository, NewSimpleService,
	)
	// nil akan diganti oleh wire melalui auto generate
	// injector can handle error add nil
	return nil, nil
}

// injector
func InitializedDatabaseRepository() *DatabaseRepository {
	wire.Build(
		NewDatabaseMysql,
		NewDatabasePostgre,
		NewDatabaseRepository,
	)
	return nil
}

// provider set untuk foo
var fooSet = wire.NewSet(NewFooRepository, NewFooService)
// provider set untuk bar
var barSet = wire.NewSet(NewBarRepository, NewBarService)

// injector foobar
func InitializedFooBarService() *FooBarService {
	// masukkan provider set
	wire.Build(fooSet, barSet, NewFooBarService)
	return nil
}