package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name      string     `gorm:"type:varchar(100);size:255;default:'' " json:"name"`
	Age       int        `gorm:"type:int;default:0" json:"age"`
	Gender    string     `gorm:"type:varchar(100);default:'male';comment:'test'" json:"gender"`
	Profile   Profile    `gorm:"ForeignKey:profile_id"`
	Languages []Language `gorm:"many2many:user_languages" json:"languages"`
}

type Profile struct {
	gorm.Model
	ProfileID  string `json:"profile_id"`
	ProfileUrl string `json:"profile_url"`
}

type Language struct {
	LanguageID int `gorm:"primary_key:language_id AUTO_INCREMENT" json:"language_id"`
	Name       string
}

func (u *User) Insert() error {
	if err := db.Create(u).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) BeforeCreate(scope *gorm.Scope) (err error) {
	scope.SetColumn("created_at", time.Now())
	return
}
func (u *User) BeforeUpdate(scope *gorm.Scope) (err error) {
	scope.SetColumn("updated_at", time.Now())
	return
}
