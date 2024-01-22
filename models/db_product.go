package models

import "time"

type Product struct {
	Id          uint      `json:"id"`
	Name        string    `json:"name" gorm:"type:varchar(50)"`
	Description string    `json:"description"`
	Price       float32   `json:"price"`
	Stock       uint      `json:"stock"`
	Date        time.Time `json:"date"`

	// Foreing Key
	Category_id   uint          `json:"category_id"`
  UserCategory  UserCategory  `json:"user_category" gorm:"foreignKey:Category_id"`
}

type Products []Product

type ProductDTO struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Stock       uint    `json:"stock"`
	Category_id uint    `json:"category_id"`
} 
