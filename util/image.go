package util

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"strings"

	"mime/multipart"
	"net/textproto"
	"os"

	"github.com/corona10/goimagehash"

	"github.com/disintegration/imaging"
	"github.com/gabriel-vasile/mimetype"
	"github.com/nfnt/resize"
	//bimg "gopkg.in/h2non/bimg.v1"
)

// const (
// 	img_diractual    = "./public/assets/images/actual"
// 	img_dirthumbnail = "./public/assets/images/thumbnail"
// 	img_dirsmall     = "./public/assets/images/small"
// )

// func Detectmime() (string, error) {
// 	return mimetype.DetectReader(file)
// }

func InitImageUpload(filePath string) {

}

func UploadPhotoForActual(width int, height int, file multipart.File, handler *multipart.FileHeader) (fileName string, pathh string, hash string, size int64, meme textproto.MIMEHeader, err error) {
	fmt.Printf("Uploading Photo: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)
	// if _, err := os.Stat(os.Getenv("Actual_Imgdir")); os.IsNotExist(err) {

	// 	os.Mkdir(os.Getenv("Actual_Imgdir"), 0755)
	// 	fmt.Println("Directory created")
	// }

	mime, err := mimetype.DetectReader(file)
	fmt.Println("Content of image", mime)

	if handler.Header.Get("Content-Type") == "image/jpeg" {

		file.Seek(0, 0)
		img, err := imaging.Decode(file) //jpeg.Decode(bytes.NewReader(data))
		if err != nil {
			fmt.Println("unable to decode jpeg: ", err)

			return "", "", "", 0, nil, err
		}

		m := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)

		fileNamee, err := os.CreateTemp(os.Getenv("Actual_Imgdir"), "gatsby-*"+handler.Filename)
		imageName := strings.TrimPrefix(fileNamee.Name(), "./public/assets/images/actual/")
		if err != nil {
			fmt.Println("Can't create temp file:", err)

			return "", "", "", 0, nil, err
		}
		fmt.Println("File path of 1st", fileNamee.Name())
		out, err := os.Create(fileNamee.Name())
		if err != nil {
			fmt.Println("Can't create file:", err)

			return "", "", "", 0, nil, err
		}
		//defer out.Close()

		fmt.Println("Congra")
		// write new image to file
		jpeg.Encode(out, m, nil)
		fi, err := os.Stat(fileNamee.Name())
		if err != nil {
			return "", "", "", 0, nil, err
		}
		hash, _ := goimagehash.AverageHash(img)

		//hashstring := fmt.Sprintf("Actual_", hash)
		return imageName, fileNamee.Name(), hash.ToString(), fi.Size(), handler.Header, err

	}
	if handler.Header.Get("Content-Type") == "image/png" {
		file.Seek(0, 0)
		img, err := imaging.Decode(file) //jpeg.Decode(bytes.NewReader(data))
		if err != nil {
			fmt.Println("unable to decode png: ", err)

			return "", "", "", 0, nil, err
		}

		m := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)

		fileNamee, err := os.CreateTemp(os.Getenv("Actual_Imgdir"), "gatsby-*"+handler.Filename)
		imageName := strings.TrimPrefix(fileNamee.Name(), "./public/assets/images/actual/")
		if err != nil {
			fmt.Println("Can't create temp file:", err)

			return "", "", "", 0, nil, err
		}
		fmt.Println("File path of 1st", fileNamee.Name())
		out, err := os.Create(fileNamee.Name())
		if err != nil {
			fmt.Println("Can't create file:", err)

			return "", "", "", 0, nil, err
		}
		//defer out.Close()

		fmt.Println("Congra")
		// write new image to file
		png.Encode(out, m)
		fi, err := os.Stat(fileNamee.Name())
		if err != nil {
			return "", "", "", 0, nil, err
		}
		hash, _ := goimagehash.AverageHash(img)

		//hashstring := fmt.Sprintf("Actual_", hash)
		return imageName, fileNamee.Name(), hash.ToString(), fi.Size(), handler.Header, err

	}
	return handler.Filename, handler.Filename, "", handler.Size, handler.Header, err
}
func UploadPhotoForThumbnail(width int, height int, file multipart.File, handler *multipart.FileHeader) (fileName string, pathh string, hash string, size int64, meme textproto.MIMEHeader, err error) {
	fmt.Printf("Uploading Photo: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)
	if _, err := os.Stat(os.Getenv("Thumbnail_Imgdir")); os.IsNotExist(err) {

		os.Mkdir(os.Getenv("Thumbnail_Imgdir"), 0755)
		fmt.Println("Directory created")
	}
	//	data, err := ioutil.ReadAll(file)
	// buff := make([]byte, 512) // docs tell that it take only first 512 bytes into consideration
	// if _, err = file.Read(buff); err != nil {
	// 	fmt.Println("Here", err) // do something with that error
	// 	return
	// }
	// buff, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// format := mime
	file.Seek(0, 0)
	mime, err := mimetype.DetectReader(file)
	fmt.Println("Content of image", mime)
	//mimestr := fmt.Sprintln(mime)
	//fmt.Println("Content of image", format)

	if handler.Header.Get("Content-Type") == "image/jpeg" {

		file.Seek(0, 0)
		img, err := imaging.Decode(file) //jpeg.Decode(bytes.NewReader(data))
		if err != nil {
			fmt.Println("unable to decode jpeg: ", err)
			//	w.WriteHeader(http.StatusInternalServerError)

			return "", "", "", 0, nil, err
		}

		m := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)

		fileNamee, err := os.CreateTemp(os.Getenv("Thumbnail_Imgdir"), "thumbnail_gatsby-*"+handler.Filename)
		if err != nil {
			fmt.Println("Can't create temp file", err)

			return "", "", "", 0, nil, err
		}
		imageName := strings.TrimPrefix(fileNamee.Name(), "./public/assets/images/thumbnail/")
		out, err := os.Create(fileNamee.Name())
		if err != nil {
			fmt.Println("Can't create file:", err)

			return "", "", "", 0, nil, err
		}
		fmt.Println("File path thumbnail", fileNamee)
		//defer out.Close()

		fmt.Println("Congra")
		// write new image to file
		out.Seek(0, 0)
		errr := jpeg.Encode(out, m, nil)
		if errr != nil {
			fmt.Println("unable to encode jpeg: ", err)
			//	w.WriteHeader(http.StatusInternalServerError)

			return "", "", "", 0, nil, err
		}
		fi, err := os.Stat(fileNamee.Name())
		if err != nil {
			return "", "", "", 0, nil, err
		}
		hash, _ := goimagehash.AverageHash(img)
		//hashstring := fmt.Sprintln(hash)
		return imageName, fileNamee.Name(), hash.ToString(), fi.Size(), handler.Header, err

	}
	if handler.Header.Get("Content-Type") == "image/png" {
		file.Seek(0, 0)
		img, err := imaging.Decode(file)
		if err != nil {
			fmt.Println("unable to decode png: ", err)

			return "", "", "", 0, nil, err
		}

		m := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)

		fileNamee, err := os.CreateTemp(os.Getenv("Thumbnail_Imgdir"), "thumbnail_gatsby-*"+handler.Filename)
		out, err := os.Create(fileNamee.Name())
		if err != nil {
			fmt.Println("Can't create file:", err)

			return "", "", "", 0, nil, err
		}
		imageName := strings.TrimPrefix(fileNamee.Name(), "./public/assets/images/thumbnail/")
		//defer out.Close()

		fmt.Println("Congra")
		// write new image to file
		png.Encode(out, m)
		fi, err := os.Stat(fileNamee.Name())
		if err != nil {
			return "", "", "", 0, nil, err
		}
		hash, _ := goimagehash.AverageHash(img)
		hashstring := fmt.Sprintln(hash)
		return imageName, fileNamee.Name(), hashstring, fi.Size(), handler.Header, err
	}
	return handler.Filename, handler.Filename, "", handler.Size, handler.Header, err
}

func UploadPhotoForSmall(width int, height int, file multipart.File, handler *multipart.FileHeader) (fileName string, pathh string, hash string, size int64, meme textproto.MIMEHeader, err error) {
	fmt.Printf("Uploading Photo: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)
	if _, err := os.Stat(os.Getenv("Small_Imgdir")); os.IsNotExist(err) {

		os.Mkdir(os.Getenv("Small_Imgdir"), 0755)
		fmt.Println("Directory created")
	}
	// data, err := ioutil.ReadAll(file)
	// buff := make([]byte, 0, len(data)) // docs tell that it take only first 512 bytes into consideration
	// if _, err = file.Read(buff); err != nil {
	// 	fmt.Println(err) // do something with that error
	// 	return
	// }
	// format := mime
	// buff, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// format := mime
	// fmt.Println("Content of image", format)
	file.Seek(0, 0)
	mime, err := mimetype.DetectReader(file)
	fmt.Println("Content of image", mime)
	//mimestr := fmt.Sprintln(mime)
	if handler.Header.Get("Content-Type") == "image/jpeg" {

		file.Seek(0, 0)
		img, err := imaging.Decode(file) //jpeg.Decode(bytes.NewReader(data))
		if err != nil {
			fmt.Println("unable to decode jpeg: ", err)

			return "", "", "", 0, nil, err
		}

		m := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)

		fileNamee, err := os.CreateTemp(os.Getenv("Small_Imgdir"), "small_gatsby-*"+handler.Filename)

		imageName := strings.TrimPrefix(fileNamee.Name(), "./public/assets/images/small/")

		out, err := os.Create(fileNamee.Name())
		if err != nil {
			fmt.Println("Can't create file:", err)

			return "", "", "", 0, nil, err
		}
		fmt.Println("File path small", fileNamee)
		// defer out.Close()

		fmt.Println("Congra")
		out.Seek(0, 0)
		// write new image to file
		errr := jpeg.Encode(out, m, nil)
		if errr != nil {
			fmt.Println("unable to encode jpeg: ", err)
			//	w.WriteHeader(http.StatusInternalServerError)

			return "", "", "", 0, nil, err
		}
		fi, err := os.Stat(fileNamee.Name())
		if err != nil {
			return "", "", "", 0, nil, err
		}
		hash, _ := goimagehash.AverageHash(img)
		//hashstring := fmt.Sprintln(hash)
		return imageName, fileNamee.Name(), hash.ToString(), fi.Size(), handler.Header, err

	}
	if handler.Header.Get("Content-Type") == "image/png" {
		file.Seek(0, 0)
		img, err := imaging.Decode(file)
		if err != nil {
			fmt.Println("unable to decode png: ", err)

			return "", "", "", 0, nil, err
		}

		m := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)

		fileNamee, err := os.CreateTemp(os.Getenv("Small_Imgdir"), "small_gatsby-*"+handler.Filename)
		imageName := strings.TrimPrefix(fileNamee.Name(), "./public/assets/images/small/")
		out, err := os.Create(fileNamee.Name())
		if err != nil {
			fmt.Println("Can't create file:", err)

			return "", "", "", 0, nil, err
		}
		// defer out.Close()

		fmt.Println("Congra 3rd")
		out.Seek(0, 0)
		// write new image to file
		png.Encode(out, m)
		fi, err := os.Stat(fileNamee.Name())
		if err != nil {
			return "", "", "", 0, nil, err
		}
		hash, _ := goimagehash.AverageHash(img)
		//hashstring := fmt.Sprintln(hash)
		return imageName, fileNamee.Name(), hash.ToString(), fi.Size(), handler.Header, err
	}
	return handler.Filename, handler.Filename, "", handler.Size, handler.Header, err
}
