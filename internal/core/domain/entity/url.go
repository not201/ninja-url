package entity

import "time"

type Url struct {
	OriginalUrl string
	ShortCode   string
	ExpiresAt   time.Time
}
