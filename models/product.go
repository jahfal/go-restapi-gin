package models

type Product struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	namaProduct string `gorm:"type:varchar(300)" json:"nama_product"`
	Deskripsi   string `gorm:"type:text" json:"deskripsi"`
}
