package model

import "time"

// URL represents a shortened URL.
type URL struct {
	Short     string    `bson:"_id"`
	Origin    string    `bson:"origin"`
	CreatedAt time.Time `bson:"createdAt"`
	Views     int       `bson:"views"`
}

// NewURL creates a new URL.
func NewURL(short, origin string) URL {
	return URL{
		Short:     short,
		Origin:    origin,
		CreatedAt: time.Now(),
	}
}
