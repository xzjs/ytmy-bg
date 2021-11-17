package model

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	GoodID uint `json:"goodid"`
	Num    int  `json:"num"`
	UserID uint `json:"userid"`
}
