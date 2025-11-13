package models

import "time"

type Toko struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	ID_User    uint      `gorm:"uniqueIndex" json:"id_user"`
	Nama_Toko  string    `gorm:"size:255" json:"nama_toko"`
	URLFoto    string    `gorm:"size:255" json:"url_foto"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`


	User *User `gorm:"foreignKey:ID_User;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user,omitempty"`
}

func (Toko) TableName() string {
	return "toko"
}