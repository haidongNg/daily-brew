package models

import (
	"daily-brew/config"
	"errors"
	"gorm.io/gorm"
	"strings"
)

type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price" grom:"decimal"`
	ImageURL    string  `json:"imageURL"`
}

// Get product
func GetProductById(id uint) (*Product, error) {
	var product Product
	if err := config.DB.Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

// Create product
func (p *Product) Create() error {
	err := config.DB.Create(&p).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return errors.New("product already exists")
		}
	}
	return err
}
