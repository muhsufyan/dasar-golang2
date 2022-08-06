# dasar-golang2
manual test

test dilakukan melalui postman
pada bagian response hasilnya masih Camel case jd harus diubah dulu (agar ketika dikirim ke user datanya bkn Camel case) sehingga dibagian data model yaitu model/web/* ditambah atribut json, tp khusus untuk cagetory_update_request.go tdk id-nya tidak perlu karena didapat dr query param

terjd error invalid connection, hal ini terjd karena kita tdk close koneksi saat memanggil rows di repository implementasi pd func FindById & FindAll

pada bagian post dan put saat name kosong error sdh bnr tp ditampilkannya lewat panic bukan error handle sehingga ke user tdk ada response apapun ini akan di perbaiki pd bagian selanjutnya yaitu restapi-20-errorhandler
NOTE ISSUE SBLMNYA BERASAL DARI HELPER PANICIFERROR DIMANA TDK ADA IF ERR != NIL JD BAGIAN INI SDH DI SOLVE