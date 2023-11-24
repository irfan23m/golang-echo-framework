package models

import (
	"echo-framework/helpers"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type M map[string]interface{}

type User struct {
	ID        uint      `json:"ID" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"not null;unique;type:varchar(191)"`
	Products  []Product `json:"products" gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Employee struct {
	ID        int    `json:"ID" form:"ID" gorm:"primaryKey" swagger:"description(ID)"`
	Full_name string `json:"full_name" form:"full_name" gorm:"not null;type:varchar(255)" swagger:"description(full_name)" valid:"required"`
	Email     string `json:"email" form:"email" gorm:"not null;unique;type:varchar(255)" swagger:"description(email)" valid:"required"`
	Password  string `json:"password" form:"password" gorm:"not null;uniqueIndex" swagger:"description(password)" valid:"required"`
	Age       int    `json:"age" form:"age" gorm:"not null;type:int" swagger:"description(age)"`
	Division  string `json:"division" form:"division" gorm:"not null;type:varchar(255)" swagger:"description(division)"`
}

type Product struct {
	ID        uint   `json:"ID" gorm:"primaryKey"`
	Name      string `json:"name" gorm:"not null;type:varchar(191)"`
	Brand     string `json:"brand" gorm:"not null;type:varchar(191)"`
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Response struct {
	ResponseCode string                 `json:"responseCode"`
	ResponseDesc string                 `json:"responseDesc"`
	Data         map[string]interface{} `json:"data"`
}

func (e *Employee) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(e)

	if errCreate != nil {
		err = errCreate
		return
	}

	e.Password = helpers.HashPass(e.Password)
	err = nil
	return
}
