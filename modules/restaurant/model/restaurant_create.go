package restaurantmodel

import (
	"strings"

	"example.com/g07-food-delivery/common"
)

type RestaurantCreate struct {
	common.SQLModel
	Name    string `json:"name" gorm:"column:name;"`
	Address string `json:"address" gorm:"column:addr;"`
}

func (RestaurantCreate) TableName() string { return Restaurant{}.TableName() }

func (data *RestaurantCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)

	if data.Name == "" {
		return ErrorNameCannotBeBlank
	}

	data.Address = strings.TrimSpace(data.Address)

	if data.Address == "" {
		return ErrorAddressCannotBeBlank
	}

	return nil
}
