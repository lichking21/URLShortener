package storage

import "time"

type URL struct {
	ID          uint64    `postgresql:"id"`
	OrigUrl     string    `postgresql:"original_url"`
	ShortCode   string    `postgresql:"short_code"`
	CreatedTime time.Time `postgresql:"created_at"`
}
