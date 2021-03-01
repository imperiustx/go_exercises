package restaurantmodel

import "github.com/imperiustx/go_excercises/common"

type RestaurantCreate struct {
	common.SQLModelCreate `json:",inline"`
	OwnerID               string        `json:"owner_id"`
	Name                  string        `json:"name"`
	Addr                  string        `json:"addr"`
	CityID                int           `json:"city_id"`
	Lat                   float64       `json:"lat"`
	Lng                   float64       `json:"lng"`
	Logo                  *common.Image `json:"logo"`
	ShippingFee           int           `json:"shipping_fee"`
	LogoImgID             int           `json:"logo_img_id" gorm:"-"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}
