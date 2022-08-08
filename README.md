# dasar-golang2
Stlh injector dibuat, kita generate injector jd kode dependency injection oleh wire

perintahnya adlh<br>
wire gen {nama_package}<br>
pd kasus kita nama_package adlh simple sehingga perintah kasus ini <br>
wire gen github.com/muhsufyan/dasar-golang2/simple<br>
atau bisa juga dg perintah <br>
wire<br>
tp syaratnya adlh hrs berada di directory file injector berada, kasus ini kita harus ada di directory simple<br>
maka akan dibuat file wire_gen.go yg mrpkn kode otomatis dependency injection (kode hsl dr auto generate)

* jd dg wire kita tdk perlu melakukan DI scra manual <br>

now if we want make simpleService tinggal panggil saja func InitializedService, ex we want make unit test (folder test)
NOTE PROGRAM TIDAK BISA MELAKUKAN AUTO GENERATE KARENA MASALAH INSTALASI WIRE DI WINDOWS, SALAH SATU SOLUSINYA DG MENAMBAHKAN FILE wire.exe kedalam root directory project ini