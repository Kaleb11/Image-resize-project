package model

import (
	"time"

	"gorm.io/gorm"
)

type Image struct {
	ID                uint64     `json:"id,omitempty" gorm:"primaryKey"`
	Name              string     `json:"name"`
	Alternative_Text  string     `json:"alternative_text"`
	Caption           string     `json:"caption"`
	Width             int64      `json:"width"`
	Height            int64      `json:"height"`
	Hash              string     `json:"hash"`
	Ext               string     `json:"ext"`
	Mime              string     `json:"mime"`
	Size              int64      `json:"size"`
	Url               string     `json:"url"`
	Preview_Url       string     `json:"preview_url"`
	Provider          string     `json:"provider"`
	Provider_MetaData string     `json:"provider_metadata"`
	Formater          []Formater `json:"formater,omitempty"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}
type Formater struct {
	ID        uint64 `json:"id,omitempty" gorm:"primaryKey"`
	Name      string `json:"name"`
	ImageID   uint64 `json:"imageID"`
	Width     int64  `json:"width"`
	Height    int64  `json:"height"`
	Hash      string `json:"hash"`
	Ext       string `json:"ext"`
	Mime      string `json:"mime"`
	Size      int64  `json:"size"`
	Path      string `json:"path"`
	Url       string `json:"url"`
	Image     *Image `json:"images,omitempty" gorm:"foreignKey:ImageID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type ImageResult struct {
	Actual_Width        int
	Small_Width         int
	Thumbnail_Width     int
	Actual_Height       int
	Small_Height        int
	Thumbnail_Height    int
	ActualImage_Name    string
	SmallImage_Name     string
	ThumbnailImage_Name string
	ActualImage_Path    string
	SmallImage_Path     string
	ThumbnailImage_Path string
	ActualImage_Hash    string
	SmallImage_Hash     string
	ThumbnailImage_Hash string
	Actual_Size         int64
	Small_Size          int64
	Thumbnail_Size      int64
	Mime                string
	Ext                 string
	Err                 interface{}
}
type ImgResult struct {
	Image_ID            uint64 `json:"Image_ID"`
	ActualImage_Path    string `json:"ActualImage_Path"`
	FormaterImage_Path1 string `json:"FormaterImage_Path1"`
	FormaterImage_Path2 string `json:"FormaterImage_Path2"`
}
