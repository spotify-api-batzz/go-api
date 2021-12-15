package models

import (
	"spotify-server/utils"
	"time"

	"github.com/google/uuid"
)

type TopArtistData struct {
	ID          uuid.UUID `db:"id" json:"id"`
	ArtistID    string    `db:"artist_id" json:"artistId"`
	TopArtistID string    `db:"top_artist_id" json:"topArtistId"`
	Order       int       `db:"order" json:"order"`

	TimePeriod string    `db:"time_period" json:"timePeriod"`
	CreatedAt  time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt  time.Time `db:"updated_at" json:"updatedAt"`
}

func NewTopArtistData(name string, artistID string, order int, timePeriod string, topArtistID string) TopArtistData {
	return TopArtistData{
		ID:          uuid.New(),
		ArtistID:    artistID,
		TopArtistID: topArtistID,
		Order:       order,
		TimePeriod:  timePeriod,
	}
}

func (r *TopArtistData) TableName() string {
	return "top_artist_data"
}

func (r *TopArtistData) TableColumns() []string {
	return []string{"id", "artist_id", `"order"`, "top_artist_id", "time_period", "created_at", "updated_at"}
}

func (a *TopArtistData) ToSlice() []interface{} {
	slice := make([]interface{}, 7)
	slice[0] = a.ID
	slice[1] = a.ArtistID
	slice[2] = a.Order
	slice[3] = a.TopArtistID
	slice[4] = a.TimePeriod
	slice[5] = utils.Now()
	slice[6] = utils.Now()

	return slice
}
