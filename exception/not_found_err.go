package exception

// struct ini mengikuti contract interface error dimana contract atribut dr interface error adlh Error()
type NotFoundError struct {
	Error string
}

// buat constructor
func NewNotFoundError(error string) NotFoundError {
	return NotFoundError{Error: error}
}
