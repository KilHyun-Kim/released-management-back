package models

type Technic struct {
	// gorm.Model
	Image    string `json:"img"`
	ImageAlt string `json:"imgalt"`
	TechId   int    `gorm:"column:tech_id" json:"tech_id"`
	TechName string `gorm:"column:tech_name" json:"tech_name"`
}

type Version struct {
	VersionId int    `gorm:"column:version_id" json:"version_id"`
	TechId    int    `gorm:"column:tech_id" json:"tech_id"`
	Title     string `json:"title"`
	SubTitle  string `json:"subtitle"`
	Contents  string `json:"contents"`
	Link      string `json:"link"`
}

// type Item struct {

// }
