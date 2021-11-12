package persistence

import (
	"image/internal/constant/model"
	"image/platform/db"

	"gorm.io/gorm/clause"
)

// ImagePersistence contains the list of functions for database table image
type ImagePersistence interface {
	AddImage(image *model.Image) (*model.Image, error)
	AddFormaterImage(formater []*model.Formater) error
	GetImageById(id uint64) (*model.Image, error)
	GetFormaterImageById(id uint64) ([]model.Formater, error)
	UpdateImage(image *model.Image, formaterImage []*model.Formater, id uint64) (*model.Image, error)
	GetFormater(searchKey string, name string) (*model.Formater, error)
	DeleteImage(id uint64) error
}

type imagePersistence struct {
	db db.DbPlatform
}

// ImageInit is to init the user persistence that contains image data
func ImageInit(db db.DbPlatform) ImagePersistence {
	return &imagePersistence{
		db,
	}
}

func (im *imagePersistence) AddImage(image *model.Image) (*model.Image, error) {
	db, err := im.db.Open()
	if err != nil {
		return nil, err
	}
	dbc, err := db.DB()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()

	err = db.Create(image).Error
	if err != nil {
		return nil, err
	}

	return image, nil
}

func (im *imagePersistence) AddFormaterImage(formater []*model.Formater) error {
	db, err := im.db.Open()
	if err != nil {
		return err
	}
	dbc, err := db.DB()
	if err != nil {
		return err
	}
	defer dbc.Close()

	err = db.Create(&formater).Error
	if err != nil {
		return err
	}

	return nil

}
func (im *imagePersistence) GetImageById(id uint64) (*model.Image, error) {
	db, err := im.db.Open()
	if err != nil {
		return nil, err
	}
	dbc, err := db.DB()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()
	image := &model.Image{}
	if err := db.Where("id = ?", id).Preload(clause.Associations).First(image).Error; err != nil {

		return nil, err
	}

	return image, nil

}
func (im *imagePersistence) GetFormaterImageById(id uint64) ([]model.Formater, error) {
	db, err := im.db.Open()
	if err != nil {
		return nil, err
	}
	dbc, err := db.DB()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()
	imageFormater := []model.Formater{}
	if err := db.Where("image_id = ?", id).Find(&imageFormater).Error; err != nil {

		return nil, err
	}

	return imageFormater, nil

}
func (im *imagePersistence) UpdateImage(image *model.Image, formaterImage []*model.Formater, id uint64) (*model.Image, error) {
	db, err := im.db.Open()
	if err != nil {
		return nil, err
	}
	dbc, err := db.DB()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return nil, err
	}
	updatedImage := *image
	err = db.First(image).Error
	if err != nil {
		return nil, err
	}

	image.ID = updatedImage.ID
	image.Name = updatedImage.Name
	image.Alternative_Text = updatedImage.Alternative_Text
	image.Caption = updatedImage.Caption
	image.Width = updatedImage.Width
	image.Height = updatedImage.Height
	image.Hash = updatedImage.Hash
	image.Ext = updatedImage.Ext
	image.Mime = updatedImage.Mime
	image.Size = updatedImage.Size
	image.Url = updatedImage.Url
	image.Preview_Url = updatedImage.Preview_Url
	image.Provider = updatedImage.Provider
	image.Provider_MetaData = updatedImage.Provider_MetaData

	if err := tx.Save(&image).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Save(&formaterImage).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	return &updatedImage, tx.Commit().Error
}

func (im *imagePersistence) GetFormater(searchKey string, name string) (*model.Formater, error) {
	db, err := im.db.Open()
	if err != nil {
		return nil, err
	}
	dbc, err := db.DB()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()

	images := model.Formater{}

	err = db.Where("name LIKE ? AND image_id = ?", "%"+searchKey+"%", name).
		Find(&images).Error

	if err != nil {

		return nil, err
	}
	return &images, nil
}
func (im *imagePersistence) DeleteImage(id uint64) error {
	db, err := im.db.Open()
	if err != nil {
		return err
	}
	dbc, err := db.DB()
	if err != nil {
		return err
	}
	defer dbc.Close()
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}
	image := &model.Image{}
	formaterImage := &model.Formater{}
	if err := tx.Where("id= ?", id).Delete(&image).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("image_id=?", id).Delete(&formaterImage).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
