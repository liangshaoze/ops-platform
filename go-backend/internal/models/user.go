package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;size:50;not null"`
	Password string `gorm:"size:100;not null"`
	Email    string `gorm:"size:100"`
	IsAdmin  bool   `gorm:"default:false"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func CreateAdminUser(db *gorm.DB) error {
	var count int64
	db.Model(&User{}).Where("username = ?", "admin").Count(&count)

	if count == 0 {
		admin := User{
			Username: "admin",
			Password: "admin123", // 首次运行会自动加密
			Email:    "admin@ops.com",
			IsAdmin:  true,
		}
		if err := db.Create(&admin).Error; err != nil {
			return err
		}
	}
	return nil
}
