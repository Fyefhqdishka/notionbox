package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: подхоядщей записи не найдено!")

type Note struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
