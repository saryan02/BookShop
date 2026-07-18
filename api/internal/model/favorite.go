package model

import "time"

type Favorite struct {
	ID        int `json:"id"`
	SessionId int
	UserID    int
	BookID    int
	CreatedAt time.Time
}
