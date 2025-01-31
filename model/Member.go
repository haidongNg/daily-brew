package model

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `json:"-"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
}

func GetMemberByEmail(db *gorm.DB, email string) (*Member, error) {
	var member Member

	if err := db.Where("email = ?", email).First(&member).Error; err != nil {
		return nil, err
	}
	return &member, nil
}

func GetMemberByID(db *gorm.DB, id int) (*Member, error) {
	var member Member
	if err := db.Where("id = ?", id).First(&member).Error; err != nil {
		return nil, err
	}
	return &member, nil
}

func (m *Member) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	m.Password = string(hashedPassword)
	return nil
}

func (m *Member) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(m.Password), []byte(password))
	return err == nil
}

func (m *Member) Save(db *gorm.DB) error {
	var member Member
	if err := db.Where("email = ?", m.Email).First(&member).Error; err == nil {
		return errors.New("member already exists")
	}
	return db.Save(m).Error
}
