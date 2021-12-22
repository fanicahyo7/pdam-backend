package model

import "time"

type Group struct {
	KodeGroup  int
	Nama       string
	Tarif1     float64
	Tarif2     float64
	Abonemen   float64
	Kompensasi float64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
