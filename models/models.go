package models

type Technic struct {
	// gorm.Model
	Image    string `json:"img"`
	ImageAlt string `json:"imgalt"`
	TechId   int    `gorm:"column:tech_id" json:"tech_id"`
	TechName string `gorm:"column:tech_name" json:"tech_name"`
}

// type Item struct {

// }
