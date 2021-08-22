package restaurantmodel

import (
	"strings"

	"example.com/g07-food-delivery/common"
)

type RestaurantUpdate struct {
	common.SQLModel
	Name    *string        `json:"name" gorm:"column:name;"`
	Address *string        `json:"address" gorm:"column:addr;"`
	Status  *int           `json:"-" gorm:"column:status;"`  // json:"-" => security
	Logo    *common.Image  `json:"logo" gorm:"column:logo;"` // Dung con tro vi muon neu khong co logo thi co the truyen null
	Cover   *common.Images `json:"cover" gorm:"column:cover;"`
}

/// Override func TabbleName cua gorm, de khi goi db.Updates no hieu la update vao table "restaurants"
/// Neu khong no se tu hieu la table name "restaurantupdates"
func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }

func (updateData *RestaurantUpdate) Validate() error {
	if strPtr := updateData.Name; strPtr != nil {
		str := strings.TrimSpace(*strPtr)
		if str == "" {
			return ErrorNameCannotBeBlank
		}
		updateData.Name = &str
	}

	if strPtr := updateData.Address; strPtr != nil {
		str := strings.TrimSpace(*strPtr)
		if str == "" {
			return ErrorAddressCannotBeBlank
		}
		updateData.Address = &str
	}

	return nil
}
