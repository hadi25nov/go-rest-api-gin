package models

type Product struct {
	id          int64  `gorm:"primaryKey" json:"id"`
	NamaProduct string `gorm:"type:varchar(300)" json:"id"`
	Deskripsi   string `gorm:"type:text" json:"id"`
}
