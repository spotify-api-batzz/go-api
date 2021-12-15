package models

import (
	"spotify-server/utils"
	"time"

	"github.com/google/uuid"
)

type RecentListenData struct {
	ID             uuid.UUID `db:"id" json:"id"`
	SongID         string    `db:"song_id" json:"songId"`
	RecentListenID string    `db:"recent_listen_id" json:"recentListenId"`
	PlayedAt       time.Time `db:"played_at" json:"playedAt"`
	CreatedAt      time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt      time.Time `db:"updated_at" json:"updatedAt"`

	Song *Song `json:"song,omitempty"`
	User *User `json:"user,omitempty"`
}

func NewRecentListenData(songID string, recentListenID string, playedAt time.Time) RecentListenData {
	return RecentListenData{
		ID:             uuid.New(),
		SongID:         songID,
		RecentListenID: recentListenID,
		PlayedAt:       playedAt,
	}
}

func (r *RecentListenData) TableName() string {
	return "recent_listen_data"
}

func (r *RecentListenData) TableColumns() []string {
	return []string{"id", "song_id", "recent_listen_id", "played_at", "created_at", "updated_at"}
}

func (r *RecentListenData) ToSlice() []interface{} {
	slice := make([]interface{}, 6)
	slice[0] = r.ID
	slice[1] = r.SongID
	slice[2] = r.RecentListenID
	slice[3] = r.PlayedAt.Format(time.RFC3339)
	slice[4] = utils.Now()
	slice[5] = utils.Now()

	return slice
}

func AttachSongsToRecentListenData(songs []Song, recentListenData []RecentListenData) []RecentListenData {
Outer:
	for i := range recentListenData {
		for _, song := range songs {
			if song.ID == recentListenData[i].SongID {
				recentListenData[i].Song = &song
				continue Outer
			}
		}
	}
	return recentListenData
}
