package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Image struct {
	Id        int    `json:"id" gorm:"column:id;"`
	Url       string `json:"url" gorm:"column:url;"`
	Width     int    `json:"width" gorm:"column:width;"`
	Height    int    `json:"height" gorm:"column:height;"`
	CloudName string `json:"cloud_name,omitempty" gorm:"-"`
	Extension string `json:"extension,omitempty" gorm:"-"`
}

func (Image) TableName() string { return "images" }

func (j *Image) Fulfill(domain string) {
	j.Url = fmt.Sprintf("%s/%s", domain, j.Url)
}

// Gorm dang khong biet lam the nao de chuyen hoa object thanh json
// Buoc phai override lai func Scan
func (j *Image) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var img Image
	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}

	*j = img
	return nil
}

//

// Value return json value, implement driver.Valuer interface
// Dung de huong dan gorm parse mot structure thanh json -> neu khong gorm se khong biet luu xuong the nao
func (j *Image) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j) // Tuong duong json.encode() o cac ngon ngu khac
}

type Images []Image

// Gorm dang khong biet lam the nao de chuyen hoa object thanh json
// Buoc phai override lai func Scan
func (j *Images) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var img []Image
	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}

	*j = img
	return nil
}

// Value return json value, implement driver.Valuer interface
// Dung de huong dan gorm parse mot structure thanh json -> neu khong gorm se khong biet luu xuong the nao
func (j *Images) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j) // Tuong duong json.encode() o cac ngon ngu khac
}
