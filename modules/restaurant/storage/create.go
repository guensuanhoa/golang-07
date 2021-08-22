package restaurantstorage

import (
	"context"

	restaurantmodel "example.com/g07-food-delivery/modules/restaurant/model"
)

// Tam thoi chap nhan luc nao cung truyen vao context, sau nay cung de traccing, 1-3 tinh nang cua golang
func (s *sqlStore) Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	db := s.db

	// Save to DB
	if err := db.Create(&data).Error; err != nil {
		// Tang nay khong phai la transpor cho nen tra ve error thoi
		// c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}
	return nil
}
