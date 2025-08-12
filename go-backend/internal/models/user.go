/*
 * @Author: 12866449444136360 liangsz@aliyun.com
 * @Date: 2025-08-12 14:01:09
 * @LastEditors: 12866449444136360 liangsz@aliyun.com
 * @LastEditTime: 2025-08-12 14:02:49
 * @FilePath: \自学项目\go-backend\internal\models\user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primarykey"`
	Username  string    `gorm:"uniqueIndex;size:50;not null"`
	Password  string    `gorm:"size:100;not null"`
	Email     string    `gorm:"size:100"`
	IsAdmin   bool      `gorm:"default:false"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// BeforeCreate 钩子 - 创建前加密密码
func (u *User) BeforeCreate(tx *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword 验证密码
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
