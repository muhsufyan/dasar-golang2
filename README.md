# dasar-golang2
BINDING INTERFACE

biasanya we use interface as kontrak struct tp wire by default use tipe data aslinya so if ada parameter yg berupa interface & tdk ada provider yg return interface tsb maka akan terjd error. Solusi agar tdk error, beri tahu wire bahwa we want melakukan binding interface (memberi tahu for a interface will use provider dg tipe apa(implementasi struct)). Artinya jika ada provider yg memerlukan dependency yg berupa interface. 

PENJELASAN MATERI INI MEMANG MEMBINGUNGKAN JD TONTON VIDEONYA 

langsung ke contoh, simple/hallo.go

kita buat juga injector-nya, tp pertama" we make injector yg salah di simple/injector.go lalu buat injector yg benar. jika ingin melihat yg salah comment injector yg benar begitupun sebaliknya

DONE