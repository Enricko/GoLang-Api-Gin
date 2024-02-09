package models

type Product struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	NamaProduct string `gorm:"varchar(300)" json:"nama_product"`
	Deskripsi   string `gorm:"text" json:"deskripsi"`
}
