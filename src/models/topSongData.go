package models

import (
	"spotify-server/utils"
	"time"

	"github.com/google/uuid"
)

type TopSongData struct {
	ID         uuid.UUID `db:"id" json:"id"`
	TopSongID  string    `db:"top_song_id" json:"topSongId"`
	SongID     string    `db:"song_id" json:"songId"`
	Order      int       `db:"order" json:"order"`
	TimePeriod string    `db:"time_period" json:"timePeriod"`
	CreatedAt  time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt  time.Time `db:"updated_at" json:"updatedAt"`

	Song *Song `json:"song,omitempty"`
	User *User `json:"user,omitempty"`
}

func (r *TopSongData) TableName() string {
	return "top_song_data"
}

func (r *TopSongData) TableColumns() []string {
	return []string{"id", "top_song_id", "song_id", `"order"`, "time_period", "created_at", "updated_at"}
}

func NewTopSongData(topSongID string, songID string, order int, timePeriod string) TopSongData {
	return TopSongData{
		ID:         uuid.New(),
		SongID:     songID,
		TopSongID:  topSongID,
		Order:      order,
		TimePeriod: timePeriod,
	}
}

func (t *TopSongData) ToSlice() []interface{} {
	slice := make([]interface{}, 7)
	slice[0] = t.ID
	slice[1] = t.TopSongID
	slice[2] = t.SongID
	slice[3] = t.Order
	slice[4] = t.TimePeriod
	slice[5] = utils.Now()
	slice[6] = utils.Now()

	return slice
}

func AttachSongsToTopSongData(topSongData []TopSongData, songs []Song) []TopSongData {
Outer:
	for i := range topSongData {
		for _, song := range songs {
			if song.ID == topSongData[i].SongID {
				topSongData[i].Song = &song
				continue Outer
			}
		}
	}
	return topSongData
}
