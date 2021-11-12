package module

import (
	"image/internal/constant/model"
	"image/internal/repository"
	"image/internal/storage/persistence"
	"image/platform/image"
	"mime/multipart"
)

var imageService *service

// UseCase contains the function of business logic of domain image
type UseCase interface {
	AddImage(image *model.Image) (*model.Image, error)
	AddFormaterImage(formater []*model.Formater) error
	AddRealImage(file multipart.File, handler *multipart.FileHeader) (Imgg *model.ImgResult, err error)
	GetImageById(id uint64) (*model.Image, error)
	GetFormaterImageById(id uint64) ([]model.Formater, error)
	UpdateRealImage(id uint64, file multipart.File, handler *multipart.FileHeader) (Imgg *model.ImgResult, err error)
	DeleteImage(id uint64) error
}

type service struct {
	imageRepo     repository.ImageRepository
	imagePersist  persistence.ImagePersistence
	imagePlatform image.ImagePlatform
}

// Initialize takes all necessary user for domain user to run the business logic of domain user
func Initialize(imageRepo repository.ImageRepository, imagePersist persistence.ImagePersistence, imagePlatform image.ImagePlatform) UseCase {
	imageService = &service{
		imageRepo:     imageRepo,
		imagePersist:  imagePersist,
		imagePlatform: imagePlatform,
	}

	return imageService
}

func ImageService() UseCase {
	return imageService
}
