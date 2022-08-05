# dasar-golang2
membuat router url

buat main.go di main directory tp sblm di main tsb memanggil/jlnkan fungsi controller maka kita perlu membuat constructor (kalau di python func __init__) dari controller, di go constructor dr func ini kita namai New....

dan contructor ini fungsinya as injeksi.

selain membuat constructor untuk controller kita juga buat constructor untuk service, & repository

di main.go isinya adlh injeksi dependensi yg diperlukan & define url 

sblmnya kita perlu juga membuat db pooling untuk koneksi ke db. 

caranya buat folder baru yaitu app/database.go yg berisi db pooling