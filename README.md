# dasar-golang2
UNIT TEST dg testify

sbnrnya ini lbh ke integration test bukan unit test (hit endpoint langsung)

Buat folder baru yaitu test

pd router (define url) akan kita pisahkan ke dlm 1 func (disimpan dlm 1 file) agar mudah saat test, di app buat file baru dan pindahkan define router(asalnya di main.go) ke sana(file baru yaitu app/router.go)

untuk test ini kita akan gunakan db test (bukan mock) jd buat dulu db baru yaitu categoryapi_test lalu migrate