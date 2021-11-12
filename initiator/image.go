package initiator

import (
	routing "image/internal/api/rest"
	"image/internal/handler/rest"
	image "image/internal/module"
	"image/internal/repository"
	"image/internal/storage/persistence"
	"image/platform/db"
	ginRouter "image/platform/gin"
	img "image/platform/image"
	"os"
)

// Image initializes the domain image
func Image(dbPlatform db.DbPlatform) []ginRouter.Router {
	actualimagePath := os.Getenv("Actual_Imgdir")
	thumbnailimagePath := os.Getenv("Thumbnail_Imgdir")
	smallimagePath := os.Getenv("Small_Imgdir")
	imageUploader := img.Initialize(actualimagePath, thumbnailimagePath, smallimagePath)

	// Initiate the image persistence
	imagePersistence := persistence.ImageInit(dbPlatform)
	// Initiate the image repository
	imageRepository := repository.ImageInit()
	// Initiate the image service
	UseCase := image.Initialize(imageRepository, imagePersistence, imageUploader)

	// Initiate the image rest API handler
	handlerr := rest.ImageInit(UseCase, imageUploader)

	// Initiate the image routing
	return routing.ImageRouting(handlerr)

}
