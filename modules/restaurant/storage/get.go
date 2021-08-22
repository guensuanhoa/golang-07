package restaurantstorage

import (
	"context"

	"example.com/g07-food-delivery/common"
	restaurantmodel "example.com/g07-food-delivery/modules/restaurant/model"
	"gorm.io/gorm"
)

func (s *sqlStore) FindDataWithCondition(
	ctx context.Context,
	cond map[string]interface{}, // Su dung nhieu noi khac nhau, vs cac condition khac nhau nen khong truyen vao model
	moreKeys ...string, // De san de sau nay khoi sua, tam chap nhan voi team la luc nao get cung co cai nay
) (*restaurantmodel.Restaurant, error) { // Tra ve con tro restaurant de truong hop error co the tra ve nil, Neu ai su dung se check duoc nil
	db := s.db

	var data restaurantmodel.Restaurant
	if err := db.Where(cond).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrDataNotFound
		}

		return nil, err
	}
	return &data, nil // Error la mot interface nen no co the ve nil duoc
}
