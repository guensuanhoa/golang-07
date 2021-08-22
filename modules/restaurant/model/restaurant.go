package restaurantmodel

import (
	"errors"

	"example.com/g07-food-delivery/common"
)

type Restaurant struct {
	common.SQLModel                /// Day goi la embed
	Name            string         `json:"name" gorm:"column:name;"`
	Address         string         `json:"address" gorm:"column:addr;"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo;"` // Dung con tro vi muon neu khong co logo thi co the truyen null
	Cover           *common.Images `json:"cover" gorm:"column:cover;"`
}

func (Restaurant) TableName() string { return "restaurants" }

var (
	ErrorNameCannotBeBlank    = errors.New("restaurant name cannot be blank")
	ErrorAddressCannotBeBlank = errors.New("restaurant address cannot be blank")
)
