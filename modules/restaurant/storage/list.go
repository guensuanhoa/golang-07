package restaurantstorage

import (
	"context"

	"example.com/g07-food-delivery/common"
	restaurantmodel "example.com/g07-food-delivery/modules/restaurant/model"
	"gorm.io/gorm"
)

func (s *sqlStore) ListDataWithCondition(
	ctx context.Context,
	paging *common.Paging,
	filter *restaurantmodel.Filter,
	moreKeys ...string, // De san de sau nay khoi sua, tam chap nhan voi team la luc nao get cung co cai nay
) ([]restaurantmodel.Restaurant, error) { // Tra ve con tro restaurant de truong hop error co the tra ve nil, Neu ai su dung se check duoc nil
	db := s.db

	var result []restaurantmodel.Restaurant

	if filter.UserId > 0 {
		db = db.Where("owner_id = ?", filter.UserId)
	}

	db = db.Where("status not in (0)")

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	offset := (paging.Page - 1) * paging.Limit

	if err := db.
		Limit(paging.Limit).
		Offset(offset).
		Order("id desc").
		Find(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrDataNotFound
		}

		return nil, err
	}
	return result, nil // Error la mot interface nen no co the ve nil duoc
}
