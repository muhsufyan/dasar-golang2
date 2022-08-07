# dasar-golang2
DEPENDENCY INJECTION
PROVIDER

sebelumnya pada restapi dibagian main.go kita telah melakukan dependency injection (DI) secara manual melalui constructor, now we use library for do DI yaitu use wire from google. nah di wire constructor as func untuk melakukan dependency injection itu disbt dg provider, jd sekarang constructor ini kita sbt as provider.

go get github.com/google/wire/cmd/wire

install dulu untuk auto generate dg perintah go install github.com/google/wire/cmd/wire@latest

lalu kita perlu menambahkan ke environment variabel lokasi dari wire tsb, untuk mengetahui lokasi GOPATH perintahnya
windows => echo %GOPATH%
mac => env | greph GOPATH

set path dari GOPATH/bin/wire

buat folder simple untuk membuat objek & provider nya (func provider). di simple/simple.go repository akan menginjek service karena service depend ke repository