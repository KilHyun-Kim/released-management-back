package dblayer

import "kilhyun-kim/released-back/models"

type DBLayer interface {
	GetAllTech() ([]models.Technic, error)
}
