package repository

// ImageRepository contains the functions of data logic for domain image
type ImageRepository interface {
}

type imageRepository struct {
}

// ImageInit initializes the data logic / repository for domain image
func ImageInit() ImageRepository {
	return imageRepository{}
}
