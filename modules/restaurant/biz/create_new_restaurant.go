package restaurantbiz

import (
	"context"

	restaurantmodel "example.com/g07-food-delivery/modules/restaurant/model"
)

// Interface trong golang thuong duoc khai bao o noi no duoc su dung
type CreateRestaurantStore interface {
	Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error
}

// Encapsulation
type createNewRestaurant struct {
	// Struct nay no chi chua 1 cai store khong tuong minh
	// Injection store tuong ung vao thi co the su dung
	// No co the la MySql, Mongo ...
	store CreateRestaurantStore
}

// O noi goi ham NewCreateRestaurantStore se biet duoc cai store nao duoc truyen vao
// Day goi la injection
func NewCreateRestaurantStore(store CreateRestaurantStore) *createNewRestaurant {
	return &createNewRestaurant{store: store}
}

func (biz *createNewRestaurant) CreateNewRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := biz.store.Create(ctx, data); err != nil {
		return err
	}

	return nil
}
