package model

import "time"

type User struct {
	NoRekening string `gorm:"primarykey"`
	Nama       string
	Alamat     string
	RT         string
	RW         string
	Telepon    string
	Foto       string
	Role       string
	Password   string
	KodeGroup  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
