# dasar-golang2
buat folder baru yaitu migrate, run perintah berikut <br>
buat database baru dg nama categoryapi<br>
migrate create -ext sql -dir migrate -seq category_table<br>
di file .up isi <br>
CREATE TABLE cateogry{
    id int AUTO_INCREMENT PRIMARY KEY,
    name varchar(200) NOT NULL
}engine=InnoDB

dan di .down isi <br>
DROP TABLE cateogry;
<br>
jlnkan migrasi ke db dg perintah<br>
migrate -path migrate -database "mysql://root:@tcp(localhost:3306)/categoryapi" up