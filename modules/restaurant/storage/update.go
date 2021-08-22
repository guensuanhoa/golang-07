package restaurantstorage

import (
	"context"

	restaurantmodel "example.com/g07-food-delivery/modules/restaurant/model"
)

func (s *sqlStore) Update(
	ctx context.Context,
	cond map[string]interface{}, // Su dung nhieu noi khac nhau, vs cac condition khac nhau nen khong truyen vao model
	updateData *restaurantmodel.RestaurantUpdate,
) error { // Tra ve con tro restaurant de truong hop error co the tra ve nil, Neu ai su dung se check duoc nil
	db := s.db

	var data restaurantmodel.Restaurant
	if err := db.Where(cond).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}
