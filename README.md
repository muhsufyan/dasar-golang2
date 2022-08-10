# dasar-golang2
CLEAN UP FUNC
maksudnya adalah setelah objek digunakan maka provider akan mengembalikan clouser func, misalnya setelah provider digunakan maka kembalikan closure func berupa close() (contoh nyatanya adalah close untuk koneksi db, close file, dll). sehingga selain return objek, error juga kita tambah return berupa closure func pd providernya.

contoh kasus ini adlh simple/file.go & simple/connection.go dimana buat file dulu baru connection, 2 file ini berisi provider func yg disertai dg clean up func berupa closure (kebetulan func closure yg direturn ke 2 file tsb pd func provider berupa func Close()).

next buat injectornya dg func InitializedConnection di simple/injector.go

kita buat juga test nya di test/file_test.go