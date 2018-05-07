package presenters

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Basabi-lab/lms/domains/models"
	"github.com/Basabi-lab/lms/usecases"
)

func TestSongGetByIDToJson(t *testing.T) {
	sgbip := NewSongGetByIDPresenter()
	song := &models.Song{
		Title: "Song title",
	}
	usecaseResult := &usecases.SongGetByIDUsecaseResult{
		Song: song,
	}

	_, err := sgbip.ToByteList(usecaseResult)
	assert.NoError(t, err)
}