package model

import (
	"database/sql"
	"time"
)

type BaseModel struct {
	CreatedAt time.Time
	UpdateAt  time.Time
	DeletedAt sql.NullTime
}
