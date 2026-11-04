package categories

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterRoutes(api huma.API, handler *CategoriesHandler) {
	huma.Register(api, huma.Operation{
		Method:      http.MethodGet,
		OperationID: "get_categories",
		Path:        "/api/v1/categories",
		Summary:     "Получить все категории",
		Tags:        []string{"Category"},
	}, handler.GetAllCategoriesHandler)
}
