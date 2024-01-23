package models

import "time"

type User struct {
	Id        uint `json:"id"`
	ProfileID uint `json:"profile_id"`
	Profile   `json:"profile"`
	Name      string `json:"name" gorm:"type:varchar(50)"`
	Mail      string `json:"mail" gorm:"type:varchar(50)"`
	Phone     string `json:"phone" gorm:"type:varchar(20)"`
	Password  string `json:"password" gorm:"type:varchar(200)"`
	Date      time.Time `json:"date"`
}

type Users []User

type UserDTO struct {
  Name      string `json:"name"`
  Mail      string `json:"mail"`
  Phone     string `json:"phone"`
  Password  string `json:"password"`
	ProfileID uint `json:"profile_id"`
}

type LoginDTO struct {
  Mail string `json:"mail"`
  Password string `json:"password"`
}

type LoginResponseDTO struct {
  Name string `json:"name"`
  Token string `json:"token"`
}
