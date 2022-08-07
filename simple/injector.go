//go:build wireinject
// +build wireinject

package simple

import "github.com/google/wire"

// buat func injector
// return-nya adlh dependency terakhir yg diperlukan (dlm kasus ini adlh service), nama func injector kita awali dg Initialize....
// ini mirip sprti Constructor (New....)
func InitializedService() *SimpleService {
	// masukkan provider apa saja yg akan diinjek/digunakan
	// INGAT FUNC MANA YG PERLU DIINJEK AKAN DIHANDLE OLEH GOOGLE WIRE, PD KASUS INI service perlu diinjek oleh repository NAH HAL INI AKAN DIHANDLE OLEH WIRE
	// JD KITA HANYA MENDAFTARKAN FUNC PROVIDER YG AKAN DIINJEK/DIGUNAKAN SAJA
	wire.Build(
		NewSimpleRepository, NewSimpleService,
	)
	// nil akan diganti oleh wire melalui auto generate
	return nil
}
