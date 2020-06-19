package domain

import (
	"time"
)

type Base struct {
	ID string `json:"user_id"; gorm:type:"uuid;primary_key" valid:"uuid"`
	CreatedAt time.Time `valid:"-"`
	UpdatedAt time.Time `valid:"-"`
}
