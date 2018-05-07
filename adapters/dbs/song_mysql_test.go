package dbs

import (
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"

	"github.com/Basabi-lab/lms/domains/models"
	"github.com/Basabi-lab/lms/test"
)

func TestSongGetAll(t *testing.T) {
	db, mock := connectDB("test_song_get_all")
	defer db.Close()

	ar := NewSongMysql(db)

	song := test.TestSongData()

	var songCols []string = []string{"id", "created_at", "updated_at", "deleted_at", "album_id", "artist_id", "title", "genre", "year", "track", "album_num", "dir"}
	mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows(songCols).
		AddRow(song.ID, song.CreatedAt, song.UpdatedAt, song.DeletedAt, song.AlbumID, song.ArtistID, song.Title, song.Genre, song.Year, song.Track, song.AlbumNum, song.Dir))

	songs, err := ar.GetAll()
	assert.NoError(t, err)

	expect := []*models.Song{}
	expect = append(expect, song)
	assert.Equal(t, expect, songs)
}

func TestSongGetByID(t *testing.T) {
	db, mock := connectDB("test_song_get_by_id")
	defer db.Close()

	ar := NewSongMysql(db)

	expect := test.TestSongData()

	var songCols []string = []string{"id", "created_at", "updated_at", "deleted_at", "album_id", "artist_id", "title", "genre", "year", "track", "album_num", "dir"}
	mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows(songCols).
		AddRow(expect.ID, expect.CreatedAt, expect.UpdatedAt, expect.DeletedAt, expect.AlbumID, expect.ArtistID, expect.Title, expect.Genre, expect.Year, expect.Track, expect.AlbumNum, expect.Dir))

	song, err := ar.GetByID(0)
	assert.NoError(t, err)

	assert.Equal(t, expect, song)
}