package restaurantbiz

import (
	"context"

	restaurantmodel "example.com/g07-food-delivery/modules/restaurant/model"
)

// Interface trong golang thuong duoc khai bao o noi no duoc su dung
type GetRestaurantStore interface {
	FindDataWithCondition(
		ctx context.Context,
		cond map[string]interface{}, // Su dung nhieu noi khac nhau, vs cac condition khac nhau nen khong truyen vao model
		moreKeys ...string, // De san de sau nay khoi sua, tam chap nhan voi team la luc nao get cung co cai nay
	) (*restaurantmodel.Restaurant, error) // Tra ve con tro restaurant de truong hop error co the tra ve nil, Neu ai su dung se check duoc nil
}

// Encapsulation
type getRestaurantBiz struct {
	// Struct nay no chi chua 1 cai store khong tuong minh
	// Injection store tuong ung vao thi co the su dung
	// No co the la MySql, Mongo ...
	store GetRestaurantStore
}

// O noi goi ham NewGetRestaurantBiz duoc cai store nao duoc truyen vao
// Day goi la injection
func NewGetRestaurantBiz(store GetRestaurantStore) *getRestaurantBiz {
	return &getRestaurantBiz{store: store}
}

func (biz *getRestaurantBiz) GetRestaurant(ctx context.Context, id int) (*restaurantmodel.Restaurant, error) {
	result, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, err
	}

	return result, nil
}
