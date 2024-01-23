package models

type Profile struct {
  Id uint `json:"id"`
  Name string `json:"name" gorm:"type:varchar(50)"`
}

type Profiles []Profile

type ProfileDTO struct {
  Name string `json:"name"`
}
