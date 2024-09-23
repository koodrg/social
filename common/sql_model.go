package common

import "time"

type SQLModel struct {
	Id        int        `json:"id" gorm:"column: id;"`
	CreatedAt *time.Time `json:"created_at,onitempty" gorm:"column: created_at;"`
	UpdatedAt *time.Time `json:"updated_at,onitempty" gorm:"column: updated_at;"`
}
