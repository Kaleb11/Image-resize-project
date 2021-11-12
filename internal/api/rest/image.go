package rest

import (
	"net/http"

	"image/internal/handler/rest"
	ginRouter "image/platform/gin"
)

// ImageRouting returns the list of routers for domain image
func ImageRouting(handler rest.ImageHandler) []ginRouter.Router {
	return []ginRouter.Router{
		{ // Upload image
			Method:  http.MethodPost,
			Path:    "/upload/image",
			Handler: handler.AddImage,
			//Middlewares: []gin.HandlerFunc{middleware.ValidateAccessToken(), middleware.ValidatePermission(constant.PermFetchUser)},
		},
		{ // Get image by id
			Method:  http.MethodGet,
			Path:    "/image/:id",
			Handler: handler.GetImageById,
			//Middlewares: []gin.HandlerFunc{middleware.ValidateAccessToken(), middleware.ValidatePermission(constant.PermFetchUser)},
		},
		{ // Get formater image by id
			Method:  http.MethodGet,
			Path:    "/image/formater/:id",
			Handler: handler.GetFormaterImageById,
			//Middlewares: []gin.HandlerFunc{middleware.ValidateAccessToken(), middleware.ValidatePermission(constant.PermFetchUser)},
		},
		{ // Get image by id
			Method:  http.MethodGet,
			Path:    "/image/pdf/:id",
			Handler: handler.GeneratePDF,
			//Middlewares: []gin.HandlerFunc{middleware.ValidateAccessToken(), middleware.ValidatePermission(constant.PermFetchUser)},
		},
		{ // Update image by id
			Method:  http.MethodPut,
			Path:    "/update/image/:id",
			Handler: handler.UpdateImage,
			//Middlewares: []gin.HandlerFunc{middleware.ValidateAccessToken(), middleware.ValidatePermission(constant.PermFetchUser)},
		},
		{ // Delte image by id
			Method:  http.MethodDelete,
			Path:    "/delete/image/:id",
			Handler: handler.DeleteImage,
			//Middlewares: []gin.HandlerFunc{middleware.ValidateAccessToken(), middleware.ValidatePermission(constant.PermFetchUser)},
		},
		// { // Search formater image
		// 	Method:  http.MethodGet,
		// 	Path:    "/image/search/:id",
		// 	Handler: handler.SearchFormater,
		// 	//Middlewares: []gin.HandlerFunc{middleware.ValidateAccessToken(), middleware.ValidatePermission(constant.PermFetchUser)},
		// },
	}
}
