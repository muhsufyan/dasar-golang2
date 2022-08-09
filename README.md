# dasar-golang2
Kadang kita perlu provider dg tipe data yg sama ini menyebabkan error. ex 2 provider tipe datanya string maka when used wire will confuse provider mana yg akan digunakan (karena ada 2 provider yg bertipe sama yaitu string)

Solusinya dg membuat alias, sehingga pd kasus 2 provider dg string maka akan menjadi provider string1 & string2

ex kita punya 2 database yaitu prostgre & mysql ke 2nya bertipe data sama yaitu Database, alih-alih memakai tipe data Database maka kita gunakan alias jd kita buat tipe data postgre & mysql yg sbnrnya adlh bertipe data Database.

ini mirip dg diawal (golang-dasar) dimana<br>
type NoKTP string<br>
type NIK string<br>


implmentasinya simple/database.go untuk yg menggunakan alias dan simple/database_err.go untuk yg tanpa alias (akan error)<br>
di simple/injector.go buat func baru untuk injector nya yaitu InitializedDatabaseRepository<br>
run perintah wire gen github.com/muhsufyan/dasar-golang2/simple<br>
hasilnya harus error, karena wire bingung ada 2 provider yg tipe datanya sama yaitu Database. errornya sesuai expected yaitu provider has multiple parameters of type *github.com/muhsufyan/dasar-golang2/simple.Databases<br>

now kita mulai ke multiple binding, pertama code di simple/database_err.go nonactivekan dulu(beri komen) lalu di simple/database.go kodenya sama sprti simple/database_err.go cuma kita tambah alias<br>

generate dg perintah wire gen github.com/muhsufyan/dasar-golang2/simple
