package helper

import (
	"github.com/muhsufyan/dasar-golang2/model/domain"
	"github.com/muhsufyan/dasar-golang2/model/web"
)

func Convert2CategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}
func Convert2CategoriesResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, Convert2CategoryResponse(category))
	}
	return categoryResponses
}
