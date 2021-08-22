package restaurantbiz

import (
	"context"

	"example.com/g07-food-delivery/common"
	restaurantmodel "example.com/g07-food-delivery/modules/restaurant/model"
)

// Interface trong golang thuong duoc khai bao o noi no duoc su dung
type ListRestaurantStore interface {
	ListDataWithCondition(
		ctx context.Context,
		paging *common.Paging,
		filter *restaurantmodel.Filter,
		moreKeys ...string, // De san de sau nay khoi sua, tam chap nhan voi team la luc nao get cung co cai nay
	) ([]restaurantmodel.Restaurant, error)
}

// Encapsulation
type listRestaurantBiz struct {
	// Struct nay no chi chua 1 cai store khong tuong minh
	// Injection store tuong ung vao thi co the su dung
	// No co the la MySql, Mongo ...
	store ListRestaurantStore
}

// O noi goi ham NewCreateRestaurantStore se biet duoc cai store nao duoc truyen vao
// Day goi la injection
func NewListRestaurantBiz(store ListRestaurantStore) *listRestaurantBiz {
	return &listRestaurantBiz{store: store}
}

func (biz *listRestaurantBiz) ListRestaurant(
	ctx context.Context,
	paging *common.Paging,
	filter *restaurantmodel.Filter,
	moreKeys ...string) ([]restaurantmodel.Restaurant, error) {
	result, err := biz.store.ListDataWithCondition(ctx, paging, filter, moreKeys...)

	if err != nil {
		return nil, err
	}

	return result, nil
}
