package image

import (
	"fmt"
	"image"
	"image/internal/constant/model"
	"image/jpeg"
	"image/png"
	"mime/multipart"
	"os"
	"strings"

	"github.com/corona10/goimagehash"
	"github.com/disintegration/imaging"
	"github.com/nfnt/resize"
)

type ImagePlatform interface {
	UploadImage(file multipart.File, handler *multipart.FileHeader) (imgResult model.ImageResult, err error)
}

type imgPlatform struct {
	Actualfilepath    string
	Thumbnailfilepath string
	Smallfilepath     string
}

func Initialize(Actualfilepath string, Thumbnailfilepath string, Smallfilepath string) ImagePlatform {
	return &imgPlatform{
		Actualfilepath:    Actualfilepath,
		Thumbnailfilepath: Thumbnailfilepath,
		Smallfilepath:     Smallfilepath,
	}
}
func fileInfo(handler *multipart.FileHeader) (string, string, string) {
	uploading := fmt.Sprintf("Uploading Photo: %+v\n", handler.Filename)
	fileSize := fmt.Sprintf("File Size: %+v\n", handler.Size)
	fileHeader := fmt.Sprintf("MIME Header: %+v\n", handler.Header)
	return uploading, fileSize, fileHeader
}
func validateImage(file multipart.File) (bool, error) {
	actualIm, _, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Println("Can't create file:", err)

		return false, err
	}
	if actualIm.Width < 600 || actualIm.Height < 600 {
		formatError := fmt.Errorf("Image width or height is less than 600")
		fmt.Println("Image width or height is less than 600")

		return false, formatError
	}
	return true, nil
}
func decodeimage(file multipart.File) (image.Image, error) {
	file.Seek(0, 0)
	img, err := imaging.Decode(file) //jpeg.Decode(bytes.NewReader(data))
	if err != nil {
		fmt.Println("unable to decode jpeg: ", err)

		return nil, err
	}
	return img, err
}
func (im *imgPlatform) pathGenerate(width int, fileName string) (*os.File, error) {
	if width == 600 {
		return os.CreateTemp(im.Actualfilepath, "gatsby-*"+fileName)
	}
	if width == 500 {
		return os.CreateTemp(im.Smallfilepath, "small_gatsby-*"+fileName)
	}
	if width == 156 {
		return os.CreateTemp(im.Thumbnailfilepath, "thumbnail_gatsby-*"+fileName)
	}
	return nil, nil
}
func imageHashGenerate(width int, img image.Image) (string, error) {

	if width == 600 {
		hash, err := goimagehash.AverageHash(img)
		hashstring := hash.ToString()
		hashed := "gatsby_" + hashstring
		return hashed, err
	}
	if width == 500 {
		hash, err := goimagehash.AverageHash(img)
		hashstring := hash.ToString()
		hashed := "small_gatsby_" + hashstring
		return hashed, err
	}
	if width == 156 {
		hash, err := goimagehash.AverageHash(img)
		hashstring := hash.ToString()
		hashed := "thumbnail_gatsby_" + hashstring
		return hashed, err
	}
	return "", nil
}
func detectImageFormat(handler *multipart.FileHeader) (string, string) {
	format := handler.Header.Get("Content-Type")
	textToTrim := "image/"
	ext := strings.TrimPrefix(format, textToTrim)
	return format, ext
}
func (im *imgPlatform) UploadImage(file multipart.File, handler *multipart.FileHeader) (imgResult model.ImageResult, err error) {
	fileInfo(handler)
	validate, err := validateImage(file)
	if !validate {
		return model.ImageResult{}, err
	}
	width := 600
	format, extension := detectImageFormat(handler)

	if format == "image/jpeg" {
		width = 600
		img, err := decodeimage(file)
		if err != nil {
			fmt.Println("unable to decode jpeg: ", err)

			return model.ImageResult{}, err
		}
		file.Seek(0, 0)
		actualResized := resize.Resize(uint(width), 0, img, resize.Lanczos3)
		actualFilepath, err := im.pathGenerate(width, handler.Filename)
		if err != nil {
			fmt.Println("Can't generate temp file:", err)

			return model.ImageResult{}, err
		}
		actualImageName := strings.TrimPrefix(actualFilepath.Name(), "./public/assets/images/actual/")
		if err != nil {
			fmt.Println("Can't create temp file:", err)

			return model.ImageResult{}, err
		}
		fmt.Println("File path of 1st", actualFilepath.Name())
		actualOut, err := os.Create(actualFilepath.Name())
		if err != nil {
			fmt.Println("Can't create file:", err)

			return model.ImageResult{}, err
		}

		actualOut.Seek(0, 0)
		jpeg.Encode(actualOut, actualResized, nil)
		actualFi, err := os.Stat(actualFilepath.Name())
		if err != nil {
			return model.ImageResult{}, err
		}
		actualHash, err := imageHashGenerate(width, img)
		if err != nil {
			return model.ImageResult{}, err
		}
		actualOut.Seek(0, 0)
		actualIm, _, err := image.DecodeConfig(actualOut)
		if err != nil {
			fmt.Println("Can't create file:", err)

			return model.ImageResult{}, err
		}
		fmt.Println("Congra 1st")
		//Small
		width = 500
		smallResized := resize.Resize(uint(width), 0, img, resize.Lanczos3)
		smallFilepath, err := im.pathGenerate(width, handler.Filename)

		smallImageName := strings.TrimPrefix(smallFilepath.Name(), "./public/assets/images/small/")
		if err != nil {
			fmt.Println("Can't create temp file:", err)

			return model.ImageResult{}, err
		}
		fmt.Println("File path of 1st", smallFilepath.Name())
		smallOut, err := os.Create(smallFilepath.Name())
		if err != nil {
			fmt.Println("Can't create file:", err)

			return model.ImageResult{}, err
		}

		smallOut.Seek(0, 0)
		jpeg.Encode(smallFilepath, smallResized, nil)
		smallFi, err := os.Stat(smallFilepath.Name())
		if err != nil {
			return model.ImageResult{}, err
		}
		smallHash, err := imageHashGenerate(width, img)
		if err != nil {
			return model.ImageResult{}, err
		}

		smallIm, _, err := image.DecodeConfig(smallOut)
		if err != nil {
			fmt.Println("Can't create file:", err)

			return model.ImageResult{}, err
		}
		fmt.Println("Congra 2nd")
		//thumbnail
		width = 156
		thumbResized := resize.Resize(uint(width), 0, img, resize.Lanczos3)
		thumbFilepath, err := im.pathGenerate(width, handler.Filename)

		thumbImageName := strings.TrimPrefix(thumbFilepath.Name(), "./public/assets/images/thumbnail/")
		if err != nil {
			fmt.Println("Can't create temp file:", err)

			return model.ImageResult{}, err
		}
		fmt.Println("File path of 1st", thumbFilepath.Name())
		thumbOut, err := os.Create(thumbFilepath.Name())
		if err != nil {
			fmt.Println("Can't create file:", err)

			return model.ImageResult{}, err
		}

		thumbOut.Seek(0, 0)
		jpeg.Encode(thumbFilepath, thumbResized, nil)
		thumbFi, err := os.Stat(thumbFilepath.Name())
		if err != nil {
			return model.ImageResult{}, err
		}
		thumbHash, err := imageHashGenerate(width, img)
		if err != nil {
			return model.ImageResult{}, err
		}

		thumbIm, _, err := image.DecodeConfig(thumbOut)
		if err != nil {
			fmt.Println("Can't decode file:", err)

			return model.ImageResult{}, err
		}
		fmt.Println("Congra 3rd")
		jpegImageResult := &model.ImageResult{
			Actual_Width:        actualIm.Width,
			Actual_Height:       actualIm.Height,
			ActualImage_Name:    actualImageName,
			ActualImage_Path:    actualFilepath.Name(),
			ActualImage_Hash:    actualHash,
			Actual_Size:         actualFi.Size(),
			Mime:                format,
			Ext:                 extension,
			Small_Width:         smallIm.Width,
			Small_Height:        smallIm.Height,
			SmallImage_Name:     smallImageName,
			SmallImage_Path:     smallFilepath.Name(),
			SmallImage_Hash:     smallHash,
			Small_Size:          smallFi.Size(),
			Thumbnail_Width:     thumbIm.Width,
			Thumbnail_Height:    thumbIm.Height,
			ThumbnailImage_Name: thumbImageName,
			ThumbnailImage_Path: thumbFilepath.Name(),
			ThumbnailImage_Hash: thumbHash,
			Thumbnail_Size:      thumbFi.Size(),
			Err:                 err,
		}
		return *jpegImageResult, err
	}
	if format == "image/png" {

		width = 600
		img, err := decodeimage(file)
		if err != nil {
			fmt.Println("unable to decode jpeg: ", err)

			return model.ImageResult{}, err
		}
		file.Seek(0, 0)
		actualResized := resize.Resize(uint(width), 0, img, resize.Lanczos3)
		actualFilepath, err := im.pathGenerate(width, handler.Filename)
		if err != nil {
			fmt.Println("Can't generate temp file:", err)

			return model.ImageResult{}, err
		}
		actualImageName := strings.TrimPrefix(actualFilepath.Name(), "./public/assets/images/actual/")
		if err != nil {
			fmt.Println("Can't create temp file:", err)

			return model.ImageResult{}, err
		}
		fmt.Println("File path of 1st", actualFilepath.Name())
		actualOut, err := os.Create(actualFilepath.Name())
		if err != nil {
			fmt.Println("Can't create file:", err)

			return model.ImageResult{}, err
		}

		actualOut.Seek(0, 0)
		png.Encode(actualOut, actualResized)
		actualFi, err := os.Stat(actualFilepath.Name())
		if err != nil {
			return model.ImageResult{}, err
		}
		actualHash, err := imageHashGenerate(width, img)
		if err != nil {
			return model.ImageResult{}, err
		}
		actualOut.Seek(0, 0)
		actualIm, _, err := image.DecodeConfig(actualOut)
		if err != nil {
			fmt.Println("Can't create file:", err)

			return model.ImageResult{}, err
		}
		fmt.Println("Congra 1st")
		//Small
		width = 500
		smallResized := resize.Resize(uint(width), 0, img, resize.Lanczos3)
		smallFilepath, err := im.pathGenerate(width, handler.Filename)

		smallImageName := strings.TrimPrefix(smallFilepath.Name(), "./public/assets/images/small/")
		if err != nil {
			fmt.Println("Can't create temp file:", err)

			return model.ImageResult{}, err
		}
		fmt.Println("File path of 1st", smallFilepath.Name())
		smallOut, err := os.Create(smallFilepath.Name())
		if err != nil {
			fmt.Println("Can't create file:", err)

			return model.ImageResult{}, err
		}

		smallOut.Seek(0, 0)
		png.Encode(smallFilepath, smallResized)
		smallFi, err := os.Stat(smallFilepath.Name())
		if err != nil {
			return model.ImageResult{}, err
		}
		smallHash, err := imageHashGenerate(width, img)
		if err != nil {
			return model.ImageResult{}, err
		}

		smallIm, _, err := image.DecodeConfig(smallOut)
		if err != nil {
			fmt.Println("Can't create file:", err)

			return model.ImageResult{}, err
		}
		fmt.Println("Congra 2nd")
		//thumbnail
		width = 156
		thumbResized := resize.Resize(uint(width), 0, img, resize.Lanczos3)
		thumbFilepath, err := im.pathGenerate(width, handler.Filename)

		thumbImageName := strings.TrimPrefix(thumbFilepath.Name(), "./public/assets/images/thumbnail/")
		if err != nil {
			fmt.Println("Can't create temp file:", err)

			return model.ImageResult{}, err
		}
		fmt.Println("File path of 1st", thumbFilepath.Name())
		thumbOut, err := os.Create(thumbFilepath.Name())
		if err != nil {
			fmt.Println("Can't create file:", err)

			return model.ImageResult{}, err
		}

		thumbOut.Seek(0, 0)
		png.Encode(thumbFilepath, thumbResized)
		thumbFi, err := os.Stat(thumbFilepath.Name())
		if err != nil {
			return model.ImageResult{}, err
		}
		thumbHash, err := imageHashGenerate(width, img)
		if err != nil {
			return model.ImageResult{}, err
		}

		thumbIm, _, err := image.DecodeConfig(thumbOut)
		if err != nil {
			fmt.Println("Can't decode file:", err)

			return model.ImageResult{}, err
		}
		fmt.Println("Congra 3rd")
		pngImageResult := &model.ImageResult{
			Actual_Width:        actualIm.Width,
			Actual_Height:       actualIm.Height,
			ActualImage_Name:    actualImageName,
			ActualImage_Path:    actualFilepath.Name(),
			ActualImage_Hash:    actualHash,
			Actual_Size:         actualFi.Size(),
			Mime:                format,
			Ext:                 extension,
			Small_Width:         smallIm.Width,
			Small_Height:        smallIm.Height,
			SmallImage_Name:     smallImageName,
			SmallImage_Path:     smallFilepath.Name(),
			SmallImage_Hash:     smallHash,
			Small_Size:          smallFi.Size(),
			Thumbnail_Width:     thumbIm.Width,
			Thumbnail_Height:    thumbIm.Height,
			ThumbnailImage_Name: thumbImageName,
			ThumbnailImage_Path: thumbFilepath.Name(),
			ThumbnailImage_Hash: thumbHash,
			Thumbnail_Size:      thumbFi.Size(),
			Err:                 err,
		}
		return *pngImageResult, err
	}
	return model.ImageResult{}, err
}
