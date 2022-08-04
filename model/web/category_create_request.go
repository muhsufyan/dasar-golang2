package web

type CategoryCreateRequest struct {
	// untuk melakukan create kita perlu request dg atribut name
	Name string `validate:"required, max=200, min=10"`
}
