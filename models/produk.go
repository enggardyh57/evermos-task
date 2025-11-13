package models

import "time"

type Produk struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	NamaProduk     string    `gorm:"size:255" json:"nama_produk"`
	Slug           string    `gorm:"size:255" json:"slug"`
	HargaReseller  string    `gorm:"size:255" json:"harga_reseller"`
	HargaKonsumen  string    `gorm:"size:255" json:"harga_konsumen"`
	Stok           uint       `json:"stok"`
	Deskripsi      string    `gorm:"type:text" json:"deskripsi"`
	IDToko         uint      `json:"id_toko"`     
	IDKategori     uint      `json:"id_kategori"`  
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`

	
	Toko     *Toko     `json:"toko" gorm:"foreignKey:IDToko;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Kategori *Kategori `json:"kategori" gorm:"foreignKey:IDKategori;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (Produk) TableName() string {
	return "produk"
}