package rest

import (
	"fmt"
	"image/internal/constant/model"
	image "image/internal/module"
	img "image/platform/image"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
	"gorm.io/gorm"
)

// ImageHandler contains the function of handler for domain Image
type ImageHandler interface {
	AddImage(c *gin.Context)
	GetImageById(c *gin.Context)
	GetFormaterImageById(c *gin.Context)
	UpdateImage(c *gin.Context)
	GeneratePDF(c *gin.Context)
	DeleteImage(c *gin.Context)
}

type imageHandler struct {
	UseCase image.UseCase
	img     img.ImagePlatform
}

// ImageInit is to initialize the rest handler for domain Image
func ImageInit(UseCase image.UseCase, img img.ImagePlatform) ImageHandler {
	return &imageHandler{
		UseCase,
		img,
	}
}

func (im *imageHandler) AddImage(c *gin.Context) {
	c.MultipartForm()

	// var image model.Image
	// var formater model.Formater

	imageFile, handler, err := c.Request.FormFile("image")
	if err != nil {
		fmt.Println("0th error", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, &model.Response{Error: err.Error()})
		return
	}
	img, err := im.UseCase.AddRealImage(imageFile, handler)
	if err != nil {

		c.AbortWithStatusJSON(http.StatusBadRequest, &model.Response{Error: err.Error()})
		return

	}

	c.JSON(http.StatusOK, &model.Response{Data: img})

}
func (im *imageHandler) UpdateImage(c *gin.Context) {
	c.MultipartForm()
	imageID := c.Param("id")
	// Convert the packageID string to uint64
	id, err := strconv.Atoi(imageID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &model.Response{Error: err.Error()})
		return
	}
	// var image model.Image
	// var formater model.Formater

	imageFile, handler, err := c.Request.FormFile("image")
	if err != nil {
		fmt.Println("0th error", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, &model.Response{Error: err.Error()})
		return
	}
	img, err := im.UseCase.UpdateRealImage(uint64(id), imageFile, handler)
	if err != nil {

		c.AbortWithStatusJSON(http.StatusBadRequest, &model.Response{Error: err.Error()})
		return

	}

	c.JSON(http.StatusOK, &model.Response{Data: img})

}
func (im *imageHandler) DeleteImage(c *gin.Context) {
	c.MultipartForm()
	imageID := c.Param("id")
	// Convert the packageID string to uint64
	id, err := strconv.Atoi(imageID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &model.Response{Error: err.Error()})
		return
	}

	errr := im.UseCase.DeleteImage(uint64(id))
	if errr != nil {

		c.AbortWithStatusJSON(http.StatusBadRequest, &model.Response{Error: err.Error()})
		return

	}

	c.JSON(http.StatusOK, &model.Response{Data: "Image " + imageID + " deleted sucessfully!!!"})

}
func (ih *imageHandler) GetImageById(c *gin.Context) {
	// Check if the diver ID param is valid
	imageID := c.Param("id")
	// Convert the imageID string to uint64
	id, err := strconv.Atoi(imageID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &model.Response{Error: err.Error()})
		return
	}

	image, err := ih.UseCase.GetImageById(uint64(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusBadRequest, &model.Response{Error: gorm.ErrRecordNotFound.Error()})
			return
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, &model.Response{Error: err.Error()})
			return
		}
	}
	c.AbortWithStatusJSON(http.StatusOK, &model.Response{Data: image})
}

// func (ih *imageHandler) SearchFormater(c *gin.Context) {
// 	// Check if the diver ID param is valid
// 	imageID := c.Param("id")
// 	// Convert the imageID string to uint64
// 	// id, err := strconv.Atoi(imageID)
// 	// if err != nil {
// 	// 	c.AbortWithStatusJSON(http.StatusBadRequest, &model.Response{Error: err.Error()})
// 	// 	return
// 	// }
// 	name := "small"
// 	image, err := ih.UseCase.SearchFormater(name, imageID)
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			c.AbortWithStatusJSON(http.StatusBadRequest, &model.Response{Error: gorm.ErrRecordNotFound.Error()})
// 			return
// 		} else {
// 			c.AbortWithStatusJSON(http.StatusBadRequest, &model.Response{Error: err.Error()})
// 			return
// 		}
// 	}
// 	c.AbortWithStatusJSON(http.StatusOK, &model.Response{Data: image})
// }
func (ih *imageHandler) GetFormaterImageById(c *gin.Context) {
	// Check if the diver ID param is valid
	imageID := c.Param("id")
	// Convert the imageID string to uint64
	id, err := strconv.Atoi(imageID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &model.Response{Error: err.Error()})
		return
	}

	image, err := ih.UseCase.GetFormaterImageById(uint64(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusBadRequest, &model.Response{Error: gorm.ErrRecordNotFound.Error()})
			return
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, &model.Response{Error: err.Error()})
			return
		}
	}
	c.AbortWithStatusJSON(http.StatusOK, &model.Response{Data: image})
}
func (ih *imageHandler) GeneratePDF(c *gin.Context) {
	// Check if the diver ID param is valid
	imageID := c.Param("id")
	// Convert the imageID string to uint64
	id, err := strconv.Atoi(imageID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &model.Response{Error: err.Error()})
		return
	}

	image, err := ih.UseCase.GetImageById(uint64(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusBadRequest, &model.Response{Error: gorm.ErrRecordNotFound.Error()})
			return
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, &model.Response{Error: err.Error()})
			return
		}
	}
	imgName := image.Name
	textToTrim := "."
	//	ttfBytes :=
	name := strings.Split(imgName, textToTrim)[0]
	nameupper := strings.ToUpper(name)
	pdf := gofpdf.New("P", "mm", "A6", os.Getenv("Fonts"))

	pdf.AddUTF8Font("FAKERECE", "B", "FAKERECE.ttf")
	//	pdf.AddFont("Receiptional_Receipt_Regular", "B", "Receiptional_Receipt_Regular.ttf")
	fontSize := 16.0
	pdf.SetFont("FAKERECE", "B", fontSize)

	// CellFormat(width, height, text, border, position after, align, fill, link, linkStr)
	pdf.SetLeftMargin(10.0)
	pdf.AddPage()

	// CellFormat(width, height, text, border, position after, align, fill, link, linkStr)
	pdf.Cell(80, 7, "Receipt page 1234")
	//func ExampleFpdf_WriteAligned()
	//ImageOptions(src, x, y, width, height, flow, options, link, linkStr)
	pdf.ImageOptions(
		image.Url,
		11.5, 20,
		50, 35,
		false,
		gofpdf.ImageOptions{ImageType: image.Ext, ReadDpi: true},
		0,
		"",
	)

	pdf.SetFont("FAKERECE", "B", 10)
	pdf.Ln(35)
	pdf.WriteAligned(0, 35, "===============================", "L")
	pdf.Ln(7)
	//pdf.SetRightMargin(50.0)
	pdf.WriteAligned(0, 35, "IMAGE NAME : "+nameupper, "")
	pdf.Ln(7)
	pdf.WriteAligned(0, 35, "IMAGE FORMAT : "+strings.ToUpper(image.Mime), "L")
	pdf.Ln(7)
	pdf.WriteAligned(0, 35, "IMAGE WIDTH : "+fmt.Sprint(image.Width), "L")
	pdf.Ln(7)
	pdf.WriteAligned(0, 35, "IMAGE HEIGHT : "+fmt.Sprint(image.Height), "L")
	pdf.Ln(7)
	pdf.WriteAligned(0, 35, "IMAGE SIZE : "+fmt.Sprint(image.Size)+" KB", "L")
	pdf.Ln(7)
	pdf.WriteAligned(0, 35, "IMAGE OWNER : "+"KALEB TIALHUN", "L")
	pdf.Ln(5)
	pdf.WriteAligned(0, 35, "===============================", "L")

	pd := pdf.OutputFileAndClose(os.Getenv("Actual_Imgdir") + "/" + name + ".pdf")
	if pd != nil {
		fmt.Println("Here is the error", pd.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, &model.Response{Error: pd.Error()})
		return

	}

	c.AbortWithStatusJSON(http.StatusOK, &model.Response{Data: pd})
}
