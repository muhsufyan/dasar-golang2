# dasar-golang2
Error

wire dpt mendeteksi error pd func Provider, otomatis return error. jd provider yg ktia buat hrs mengembalikan 2 data dimana data ke 2 adlh error. 

lbh detail lihat simple/simple.go

tambah if kondisi untuk handle provider error pd constructor provider

agar diinjector-nya dpt mengembalikan error juga maka kita perlu menambahkan error di injectornya (simple/injector.go)

stlh itu generate lagi dg perintah<br>
wire gen github.com/muhsufyan/dasar-golang2/simple

di test/simple_service_test.go tambahkan juga error

now coba jika providernya mengambalikan error dg set simple/simple.go dimana Provider repository (NewSimpleRepository) return Errornya adlh true