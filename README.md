# dasar-golang2
INJECTOR

after make provider next make injector

injector is func constructor contain config to wire (maksudnya func apa saja yg akan diinjeksi kedalam func constructor)

injector ini adlh file yg akan di generate oleh google wire menjadi file wire_.go

untuk membuat injector maka pd file nya perlu ditambahkan code berikut di bagian paling atas<br>
//go:build wireinject
// +build wireinject

file injector ini disarankan dibuat secara khusus yg hanya berisi func" injector saja (tdk ada program yg lain selain func injector)

simple/injector.go akan menjd func injector