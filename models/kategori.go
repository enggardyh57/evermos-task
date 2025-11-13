package models

import "time"

type Kategori struct {
	ID         	   uint      `gorm:"primaryKey" json:"id"`
	NamaCategory   string    `gorm:"size:255;unique" json:"nama_category"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (Kategori) TableName() string {
	return "kategori"
}