# LOGGING 
ENTRY

each log yg we send maka will  made object entry.

ex when we make formatter kita sendiri maka parameter untuk melakukan formatting bukan string message but object entry (type param is object entry)

look entry https://github.com/sirupsen/logrus/blob/master/entry.go that contain data, field, level, message, etc so all log wrap by entry

for make entry NewEntry()

membuat entry scra manual sangat jarang dilakukan, ini hanya belajar supaya kita tahu saja
