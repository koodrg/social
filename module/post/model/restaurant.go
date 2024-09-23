package restaurantmodel

import (
	"social/common"

	"gorm.io/datatypes"
)

type Restaurant struct {
	common.SQLModel
	OwnerId           uint32         `json:"owner_id" gorm:"column:owner_id"`
	Name              string         `json:"name" gorm:"column:name"`
	Address           string         `json:"address" gorm:"column:addr"`
	CityId            uint32         `json:"city_id" gorm:"column:cityt_id"`
	Lat               float64        `json:"lat" gorm:"column:lat"`
	Lng               float64        `json:"lng" gorm:"column:lng"`
	Cover             datatypes.JSON `json:"cover" gorm:"column:cover"`
	Logo              datatypes.JSON `json:"logo" gorm:"column:logo"`
	Thumbnail         string         `json:"thumbnail" gorm:"column:thumbnail"`
	Status            int            `json:"status" gorm:"column:status"`
	ShippingFreePerKm float64        `json:"shipping_free_per_km" gorm:"ship_free_per_km"`
}

func (Restaurant) TableName() string { return "restaurants" }

type RestaurantCreate struct {
	common.SQLModel
	OwnerId           uint32         `json:"owner_id" gorm:"column:owner_id"`
	Name              string         `json:"name" gorm:"column:name"`
	Address           string         `json:"address" gorm:"column:addr"`
	CityId            uint32         `json:"city_id" gorm:"column:cityt_id"`
	Lat               float64        `json:"lat" gorm:"column:lat"`
	Lng               float64        `json:"lng" gorm:"column:lng"`
	Cover             datatypes.JSON `json:"cover,onitempty" gorm:"column:cover"`
	Logo              datatypes.JSON `json:"logo,onitempty" gorm:"column:logo"`
	Status            int            `json:"status" gorm:"column:status"`
	ShippingFreePerKm float64        `json:"shipping_free_per_km" gorm:"ship_free_per_km"`
}

func (RestaurantCreate) TableName() string { return Restaurant{}.TableName() }

type RestaurantUpdate struct {
	common.SQLModel
	OwnerId           uint32         `json:"owner_id" gorm:"column:owner_id"`
	Name              string         `json:"name" gorm:"column:name"`
	Address           string         `json:"address" gorm:"column:addr"`
	CityId            uint32         `json:"city_id" gorm:"column:cityt_id"`
	Lat               float64        `json:"lat" gorm:"column:lat"`
	Lng               float64        `json:"lng" gorm:"column:lng"`
	Cover             datatypes.JSON `json:"cover" gorm:"column:cover"`
	Logo              datatypes.JSON `json:"logo" gorm:"column:logo"`
	Thumbnail         string         `json:"thumbnail" gorm:"column:thumbnail"`
	Status            int            `json:"status" gorm:"column:status"`
	ShippingFreePerKm float64        `json:"shipping_free_per_km" gorm:"ship_free_per_km"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }

func (r *Restaurant) Mask() {

}
