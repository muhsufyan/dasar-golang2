# dasar-golang2
dibagian imp service blm ada validasi data input, kita akan validasi di service layer menggunakan go-playground/validator

install go get github.com/go-playground/validator/v10

lihat dokumentasinya https://github.com/go-playground/validator

yg divalidasi di struct-nya (data model) dg menambahkan validate:syarat-syarat validasi

bagian yg divalidasi adlh service create & update, untuk lbh lanjut lihat model/web/category_create_request.go untuk service create dan model/web/category_update_request.go untuk service update 

lalu dibagian service create dan update injek validasi data input-nya melalui struct impl