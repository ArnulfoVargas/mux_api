package models

type UserCategory struct {
  Id uint
  Name string `gorm:"type:varchar(50)" json:"name"`
  Message string `gorm:"type:varchar(100)" json:"message"`
}
