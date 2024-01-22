package models

type Picture struct {
  Id int `json:"id"`
  Name string `json:"name" gorm:"type:varchar(50)"`
  ProductId uint `json:"product_id"`
  Product Product `json:"product" gorm:"foreignKey:ProductId"`
}

type Pictures []Picture
