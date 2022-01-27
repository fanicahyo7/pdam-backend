package model

import "time"

type Group struct {
	Kode       string `gorm:"primarykey" binding:"required"`
	Nama       string `binding:"required"`
	Tarif1     float64
	Tarif2     float64
	Abonemen   float64
	Kompensasi float64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
