package model

import (
	"log"

	"gorm.io/gorm"
)

type Goods struct {
	gorm.Model
	Title  string
	UserID uint // foreign (default format)
}

type GoodsModel struct {
	DB *gorm.DB
}

func (gm *GoodsModel) Insert(newItem Goods) (Goods, error) {
	err := gm.DB.Create(&newItem).Error
	if err != nil {
		log.Println("insert goods query error : ", err.Error())
		return Goods{}, err
	}

	return newItem, nil
}

func (gm *GoodsModel) GetAllByID(id int) ([]Goods, error) {
	res := []Goods{}
	err := gm.DB.Where("user_id = ?", id).Find(&res).Error
	if err != nil {
		log.Println("select goods query error : ", err.Error())
		return nil, err
	}

	return res, nil
}
