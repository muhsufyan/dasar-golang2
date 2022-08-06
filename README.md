# dasar-golang2
error Handler

sblmnya error dihandle melalui panic tp ke user tdk ada apa response apapun, hal tsb mash lbh baik drpdditampilkan errornya ke user.

buat folder dg nama exception untuk handle error

ganti pada implementasi service dari panic error jd error handler(response error)

untuk handle error validasi kita akan menggunakan "class" dari library validator

NOTE : KODE SEBLMNYA MSH TERDPT ERROR YANG DIAKIBATKAN KESALAHAN PENEMAPATAN SPRTI HRSNYA POINTER INI MALAH TDK ATAU SEBALIKNYA, SEHARUSNYA DISIMPAN DIBAGIAN CONTROLLER MALAH DISIMPAN DI SERVICE