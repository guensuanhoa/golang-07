package restaurantbiz

import (
	"context"
	"errors"
	"fmt"

	"example.com/g07-food-delivery/common"
	restaurantmodel "example.com/g07-food-delivery/modules/restaurant/model"
)

// Interface trong golang thuong duoc khai bao o noi no duoc su dung
type DeleteRestaurantStore interface {
	// Ben get_restaurant.go cung co phuong thuc nay
	// Nhung khong duoc goi truc tiep qua, vi ben do se co business rieng de query
	// Cho nen phai implement lai
	FindDataWithCondition(
		ctx context.Context,
		cond map[string]interface{}, // Su dung nhieu noi khac nhau, vs cac condition khac nhau nen khong truyen vao model
		moreKeys ...string, // De san de sau nay khoi sua, tam chap nhan voi team la luc nao get cung co cai nay
	) (*restaurantmodel.Restaurant, error) // Tra ve con tro restaurant de truong hop error co the tra ve nil, Neu ai su dung se check duoc nil

	Update(
		ctx context.Context,
		cond map[string]interface{}, // Su dung nhieu noi khac nhau, vs cac condition khac nhau nen khong truyen vao model
		updateData *restaurantmodel.RestaurantUpdate,
	) error // Tra ve con tro restaurant de truong hop error co the tra ve nil, Neu ai su dung se check duoc nil
}

// Encapsulation
type deleteRestaurantBiz struct {
	// Struct nay no chi chua 1 cai store khong tuong minh
	// Injection store tuong ung vao thi co the su dung
	// No co the la MySql, Mongo ...
	store DeleteRestaurantStore
}

// O noi goi ham NewGetRestaurantBiz duoc cai store nao duoc truyen vao
// Day goi la injection
func NewDeleteRestaurantBiz(store UpdateRestaurantStore) deleteRestaurantBiz {
	return deleteRestaurantBiz{store: store}
}

func (biz deleteRestaurantBiz) DeleteRestaurant(ctx context.Context, id int) error {
	// Day la trach nhiem cua business logic
	// Find data
	oldData, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})

	fmt.Printf("OldData: %v", oldData)

	// Check loi
	if err != nil {
		if err == common.ErrDataNotFound {
			return errors.New("data not found")
		}
		return err
	}

	if oldData.Status == 0 {
		return errors.New("data has been deleted")
	}

	zero := 0

	// Va goi tang storage update
	if err := biz.store.Update(ctx, map[string]interface{}{"id": id}, &restaurantmodel.RestaurantUpdate{Status: &zero}); err != nil {
		return err
	}

	return nil
}
