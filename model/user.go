package model

type User struct {
	KodePelanggan string `gorm:"primarykey"`
	NamaPelanggan string
}
