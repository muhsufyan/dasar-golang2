package simple

type Database struct {
	Name string
}

// buat alias
type DatabasePostgre Database
type DatabaseMysql Database

// buat constructor (provider). langsung hardcode
func NewDatabasePostgre() *DatabasePostgre {
	// paksa konversi jd DatabasePostgre
	return (*DatabasePostgre)(&Database{Name: "Postgre"})
}

// buat constructor (provider). langsung hardcode
func NewDatabaseMysql() *DatabaseMysql {
	// paksa konversi jd DatabaseMysql
	return (*DatabaseMysql)(&Database{Name: "Mysql"})
}

type DatabaseRepository struct {
	DatabasePostgre *DatabasePostgre
	DatabaseMysql   *DatabaseMysql
}

// buat constructor (provider) dg tipe data yg sama yaitu database
func NewDatabaseRepository(postgre *DatabasePostgre, mysql *DatabaseMysql) *DatabaseRepository {
	return &DatabaseRepository{DatabasePostgre: postgre, DatabaseMysql: mysql}
}
