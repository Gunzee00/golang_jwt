package models

type User struct {
	Id          uint   `gorm:"primaryKey" json:"id"` // Ganti string dengan uint untuk auto increment
	NamaLengkap string `gorm:"type:varchar(300)" json:"nama_lengkap"`
	Username    string `gorm:"type:varchar(300);unique" json:"username"` // username harus unique
	Password    string `gorm:"type:varchar(300)" json:"password"`
}
