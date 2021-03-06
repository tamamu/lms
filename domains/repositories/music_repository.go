package repositories

import "github.com/Basabi-lab/lms/domains/models"

type MusicRepository interface {
	GetByID(id int64) (*models.Music, error)
	Create(music *models.Music) (int64, error)
}
