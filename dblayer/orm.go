package dblayer

import (
	"kilhyun-kim/released-back/models"

	"github.com/jinzhu/gorm"
)

type DBORM struct {
	*gorm.DB
}

func NewORM(dbname, con string) (*DBORM, error) {
	db, err := gorm.Open(dbname, con)
	return &DBORM{
		DB: db,
	}, err
}

func (db *DBORM) GetAllTech() (tech []models.Technic, err error) {
	return tech, db.Find(&tech).Error
}

func (db *DBORM) FindVersionId(id int) (version []models.Version, err error) {
	return version, db.Table("versions").Select("*").Where("tech_id=?", id).Scan(&version).Error
	// return
}
