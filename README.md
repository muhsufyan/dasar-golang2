# dasar-golang2
Untuk melihat dokumentasi api-nya kunjungi https://editor.swagger.io/

dokumentasi security untuk login

untuk login (otentifikasi) nya menggunakan api-key

untuk menambahkannya buat dulu component baru (securitySchemes) dan beri nama (kasus ini CategoryAuth) lalu terapkan di setiap http method stlh get misalnya tambah 
    security:
        - CategoryAuth: []