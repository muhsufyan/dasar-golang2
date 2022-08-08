# dasar-golang2
INJECTOR PARAMETER

parameter dalam injector, sblmnya tdk ada parameter sekarang we add parameter.

kita jdkan error as parameter jd tdk hardcode lagi (simple/simple.go) lalu di simple/injector.go tambah isError as provider parameter, sbnrnya bisa juga membuat provider lagi untuk isError (isError as provider) tp sekarang kita sdng bljr injecotr parameter.

jlnkan perintah wire gen github.com/muhsufyan/dasar-golang2/simple<br>

di test/simple_service_test.go kita ubah dg menambahkan injector parameter. jd injector parameter di injector.go dijdkan as provider parameter. Diambil dari tipe datanya (pd kasus ini bool) bukan nama variabelnya (pd kasus ini isError)