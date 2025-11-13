package models

import "time"

type User struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Nama           string    `gorm:"size:255" json:"nama"`
	Kata_sandi     string    `gorm:"size:255" json:"kata_sandi"`
	Notelp         string    `gorm:"size:255;unique" json:"notelp"`
	Tanggal_Lahir  string    `gorm:"type:date" json:"tanggal_lahir"`
	Jenis_Kelamin  string    `gorm:"size:255" json:"jenis_kelamin"`
	Tentang        string    `gorm:"type:text" json:"tentang"`
	Pekerjaan      string    `gorm:"size:255" json:"pekerjaan"`
	Email         string    `gorm:"size:255;unique" json:"email"`
	ID_Provinsi    string    `gorm:"size:255" json:"id_provinsi"`
	ID_Kota        string    `gorm:"size:255" json:"id_kota"`
	IsAdmin 	   bool      `gorm:"default:false" json:"is_admin"`
	CreatedAt 	   time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`

	
	Toko *Toko `gorm:"foreignKey:IDUser;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"toko,omitempty"`
}

func (User) TableName() string {
	return "user"
}