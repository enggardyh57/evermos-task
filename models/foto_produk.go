package models

import "time"

type Foto_Produk struct {
	ID         	   uint      `gorm:"primaryKey" json:"id"`
	ID_Produk uint      `json:"id_produk,omitempty"`
	URLFoto        string    `gorm:"size:255" json:"url_foto"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (Foto_Produk) TableName() string {
	return "foto_produk"
}