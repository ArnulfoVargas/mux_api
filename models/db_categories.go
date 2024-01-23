package models

type UserCategory struct {
	Id      uint   `json:"id"`
	Name    string `gorm:"type:varchar(50)" json:"name"`
	Message string `gorm:"type:varchar(100)" json:"message"`
}

type UserCategories []UserCategory

type UserCategoryDTO struct {
	Name    string `gorm:"type:varchar(50)" json:"name"`
	Message string `gorm:"type:varchar(100)" json:"message"`
}

