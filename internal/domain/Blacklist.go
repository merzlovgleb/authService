package domain

import "time"

type Blacklist struct {
	ID          int       `json:"id"`
	AccessToken string    `json:"access_token"`
	CreatedAt   time.Time `json:"created_at"`
}
