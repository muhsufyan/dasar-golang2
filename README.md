# dasar-golang2
STRUCT FIELD PROVIDER

wire support make provider from field sebuah provider, ex we have data dlm struct & we want menjdkan field dlm struct tsb as provider.

artinya suatu field dr struct dpt dijdkan DI untuk provider lainnya

contoh kasusnya we make file simple/configuration.go lalu buat injector-nya di injector.go dg nama InitializedConfiguration