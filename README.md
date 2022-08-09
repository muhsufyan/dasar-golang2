# dasar-golang2
PROVIDER SET

used for grouping provider. terapkan provider set saat provider nya sudah sangat banyak.

ex at simple directory make file foo.go bar.go dari 2 file tsb we have 4 provider yaitu NewFooRepository, NewFooService, NewBarRepository, NewBarService<br>
next we make new file in simple directory foobar.go, FooBarService will accept provider FooService & BarService. so now we have 5 provider (1 provider is FooBarService)

di simple/injector.go (file injector) kita akan buat grouping provider menggunakan provider set, tp sbnrnya kita bisa saja langsung membuat kode generate-nya tanpa provider set melalui wire.Build sprti sblm"nya tp agar kode lbh rapih kita akan menggunakan provider set.

Semua provider foo akan disimpan dlm variabel fooSet (grouping provider foo) & Semua provider bar akan disimpan dlm variabel barSet (grouping provider bar).<br>
di fooSet & barSet kita buat provider set untuk foo dan untuk bar. kemudian buat injector untuk foobar

lalu generate dg perintah wire gen github.com/muhsufyan/dasar-golang2/simple

NOTE KODE PROGRAM INI BELUM SELESAI, INJECTOR BELUM DIGENERATE