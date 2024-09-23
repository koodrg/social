package restaurantbusiness

import (
	"context"
	restaurantmodel "social/module/post/model"
)

type CreateRestaurantStore interface {
	CreateRestaurant(context context.Context, data *restaurantmodel.RestaurantCreate) error
}

type createRestaurantBusiness struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBusiness(store CreateRestaurantStore) *createRestaurantBusiness {
	return &createRestaurantBusiness{store: store}
}

func (biz *createRestaurantBusiness) CreateNewRestaurant(context context.Context, data *restaurantmodel.RestaurantCreate) error {
	//logic here

	if err := biz.store.CreateRestaurant(context, data); err != nil {
		return err
	}

	return nil
}
