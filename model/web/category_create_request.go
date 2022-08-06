package web

type CategoryCreateRequest struct {
	// untuk melakukan create kita perlu request dg atribut name
	Name string `validate:"required,min=1,max=100" json:"name"`
}
