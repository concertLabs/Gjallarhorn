package model

import (
	"time"
)

type Notiz struct {
	ID        int
	LiedID    int
	Text      string
	CreatedAt time.Time
}
