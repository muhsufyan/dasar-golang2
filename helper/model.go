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
