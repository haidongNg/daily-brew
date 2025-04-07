package models

import (
	"daily-brew/config"
	"errors"
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `json:"-"`
	Phone    string `json:"phone"`
	Role     string `json:"role"` // admin, staff, customer
}

func GetMemberByEmail(email string) (*Member, error) {
	var member Member

	if err := config.DB.Where("email = ?", email).First(&member).Error; err != nil {
		return nil, err
	}
	return &member, nil
}

func GetMemberByID(id uint) (*Member, error) {
	var member Member
	if err := config.DB.Where("id = ?", id).First(&member).Error; err != nil {
		return nil, err
	}
	return &member, nil
}

func (m *Member) Save() error {
	var member Member
	if err := config.DB.Where("email = ?", m.Email).First(&member).Error; err == nil {
		return errors.New("member already exists")
	}
	return config.DB.Save(m).Error
}
