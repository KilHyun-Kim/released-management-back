package dblayer

import "kilhyun-kim/released-back/models"

type DBLayer interface {
	GetAllTech() ([]models.Technic, error)
	FindVersionId(int) ([]models.Version, error)
}
