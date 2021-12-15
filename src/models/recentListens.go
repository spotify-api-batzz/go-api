package models

import (
	"spotify-server/utils"
	"time"

	"github.com/google/uuid"
)

type RecentListen struct {
	ID     uuid.UUID `db:"id" json:"id"`
	UserID string    `db:"user_id" json:"userId"`

	CreatedAt time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
}

func NewRecentListen(userID string) RecentListen {
	return RecentListen{
		ID:     uuid.New(),
		UserID: userID,
	}
}

func (r *RecentListen) TableName() string {
	return "recent_listens"
}

func (r *RecentListen) TableColumns() []string {
	return []string{"id", "user_id", "created_at", "updated_at"}
}

func (r *RecentListen) ToSlice() []interface{} {
	slice := make([]interface{}, 4)
	slice[0] = r.ID
	slice[1] = r.UserID
	slice[2] = utils.Now()
	slice[3] = utils.Now()

	return slice
}
