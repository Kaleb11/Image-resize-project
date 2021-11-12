package module

import (
	"fmt"
	"image/internal/constant/model"
	"mime/multipart"
	"strconv"

	"gorm.io/gorm"
)

func (s *service) AddImage(image *model.Image) (*model.Image, error) {
	return s.imagePersist.AddImage(image)
}

func (s *service) AddFormaterImage(formater []*model.Formater) error {
	return s.imagePersist.AddFormaterImage(formater)
}
func (s *service) GetImageById(id uint64) (*model.Image, error) {
	return s.imagePersist.GetImageById(id)
}
func (s *service) GetFormaterImageById(id uint64) ([]model.Formater, error) {
	return s.imagePersist.GetFormaterImageById(id)
}
func (s *service) DeleteImage(id uint64) error {
	return s.imagePersist.DeleteImage(id)
}

func (s *service) AddRealImage(file multipart.File, handler *multipart.FileHeader) (Imgg *model.ImgResult, err error) {
	// var image *model.Image
	// var formater *model.Formater
	img, err := s.imagePlatform.UploadImage(file, handler)
	if err != nil {
		fmt.Println("1st error", err)

		return nil, err
	}

	image := &model.Image{}
	image.Name = img.ActualImage_Name
	image.Alternative_Text = ""
	image.Caption = ""
	image.Width = int64(img.Actual_Width)
	image.Height = int64(img.Actual_Height)
	image.Hash = img.ActualImage_Hash
	image.Ext = img.Ext
	image.Mime = img.Mime
	image.Size = img.Actual_Size
	image.Url = img.ActualImage_Path
	image.Preview_Url = ""
	image.Provider = "local"
	image.Provider_MetaData = ""

	addimg, err := s.imagePersist.AddImage(image)
	if err != nil {
		if err == gorm.ErrRecordNotFound {

			return nil, err
		} else {
			return nil, err
		}
	}

	formater := []*model.Formater{}

	formater1 := &model.Formater{
		Name:    img.ThumbnailImage_Name,
		ImageID: addimg.ID,
		Width:   int64(img.Thumbnail_Width),
		Height:  int64(img.Thumbnail_Height),
		Hash:    img.ThumbnailImage_Hash,
		Size:    img.Thumbnail_Size,
		Mime:    image.Mime,
		Ext:     image.Ext,
		Path:    img.ThumbnailImage_Path,
		Url:     img.ThumbnailImage_Path,
	}
	formater2 := &model.Formater{
		Name:    img.SmallImage_Name,
		ImageID: addimg.ID,
		Width:   int64(img.Small_Width),
		Height:  int64(img.Small_Height),
		Hash:    img.SmallImage_Hash,
		Size:    img.Small_Size,
		Mime:    image.Mime,
		Ext:     image.Ext,
		Path:    img.SmallImage_Path,
		Url:     img.SmallImage_Path,
	}
	formater = append(formater, formater1, formater2)
	errr := s.imagePersist.AddFormaterImage(formater)
	if errr != nil {

		return nil, err
	}
	imgResult := &model.ImgResult{
		Image_ID:            image.ID,
		ActualImage_Path:    image.Url,
		FormaterImage_Path1: formater1.Url,
		FormaterImage_Path2: formater2.Url,
	}
	return imgResult, err
}
func (s *service) UpdateRealImage(id uint64, file multipart.File, handler *multipart.FileHeader) (Imgg *model.ImgResult, err error) {
	// var image *model.Image
	// var formater *model.Formater
	img, err := s.imagePlatform.UploadImage(file, handler)
	if err != nil {
		fmt.Println("1st error", err)

		return nil, err
	}

	image := &model.Image{}
	image.ID = id
	image.Name = img.ActualImage_Name
	image.Alternative_Text = ""
	image.Caption = ""
	image.Width = int64(img.Actual_Width)
	image.Height = int64(img.Actual_Height)
	image.Hash = img.ActualImage_Hash
	image.Ext = img.Ext
	image.Mime = img.Mime
	image.Size = img.Actual_Size
	image.Url = img.ActualImage_Path
	image.Preview_Url = ""
	image.Provider = "local"
	image.Provider_MetaData = ""

	formater := []*model.Formater{}
	smallName := "small"
	thumbName := "thumb"
	idConv := strconv.Itoa(int(id))
	formaterOne, err := s.imagePersist.GetFormater(smallName, idConv)
	if err != nil {
		fmt.Println("Can't get formater one", err)

		return nil, err
	}
	formaterTwo, err := s.imagePersist.GetFormater(thumbName, idConv)
	if err != nil {
		fmt.Println("Can't get formater one", err)

		return nil, err
	}

	formater1 := &model.Formater{
		ID:      formaterOne.ID,
		Name:    img.ThumbnailImage_Name,
		ImageID: id,
		Width:   int64(img.Thumbnail_Width),
		Height:  int64(img.Thumbnail_Height),
		Hash:    img.ThumbnailImage_Hash,
		Size:    img.Thumbnail_Size,
		Mime:    image.Mime,
		Ext:     image.Ext,
		Path:    img.ThumbnailImage_Path,
		Url:     img.ThumbnailImage_Path,
	}
	formater2 := &model.Formater{
		ID:      formaterTwo.ID,
		Name:    img.SmallImage_Name,
		ImageID: id,
		Width:   int64(img.Small_Width),
		Height:  int64(img.Small_Height),
		Hash:    img.SmallImage_Hash,
		Size:    img.Small_Size,
		Mime:    image.Mime,
		Ext:     image.Ext,
		Path:    img.SmallImage_Path,
		Url:     img.SmallImage_Path,
	}
	formater = append(formater, formater1, formater2)
	_, errr := s.imagePersist.UpdateImage(image, formater, id)
	if errr != nil {

		return nil, err
	}
	imgResult := &model.ImgResult{
		Image_ID:            image.ID,
		ActualImage_Path:    image.Url,
		FormaterImage_Path1: formater1.Url,
		FormaterImage_Path2: formater2.Url,
	}
	return imgResult, err
}
