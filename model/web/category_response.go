package web

// data model ini ini akan menjadi data response ke user
type CategoryResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
