# LOGGING 
LEVEL

level : penentuan prioritas / jenis dr suatu kejadian

semakin tinggi level maka semakin fatal, error, penting.

kita dpt menentukan level

menit 1:00 jelaskan tingkatan level, semakin kebawah jenis level semakin fatal, penting

yg dicetak (print) hanya dari level Info sampai kebawah/terakhir (Panic) jd level trace & debug tdk diprint. tp kita bisa set level apa saja yg boleh diprint(lbh tepatnya mulai dari mana karen by default mulai dari info), contoh semua level / mulai dari level error, dst caranya adlh dg set logger.SetLevel()

diserver cukup dari info saja karena akan terlalu banyak jiga mulai dari trace
