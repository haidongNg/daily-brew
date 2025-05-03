package models

import (
	"daily-brew/config"
	"errors"
	"gorm.io/gorm"
	"strings"
)

type Member struct {
	gorm.Model
	FullName string `json:"fullName"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `json:"-"`
	Phone    string `json:"phone"`
	Role     string `json:"role"` // admin, staff, customer
}

func (m *Member) GetMemberByEmail() error {
	if err := config.DB.Where("email = ?", m.Email).First(&m).Error; err != nil {
		return err
	}
	return nil
}

func (m *Member) GetMemberByID(id uint) error {
	if err := config.DB.Where("id = ?", id).First(&m).Error; err != nil {
		return err
	}
	return nil
}

func (m *Member) Create() error {
	err := config.DB.Create(&m).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return errors.New("member already exists")
		}
	}
	return err
}

func (m *Member) Update() error {
	return config.DB.Save(&m).Error
}
