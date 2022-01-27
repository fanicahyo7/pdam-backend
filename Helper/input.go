package helper

type Inputgroup struct {
	Nama       string  `binding:"required"`
	Tarif1     float64 `binding:"required"`
	Tarif2     float64 `binding:"required"`
	Abonemen   float64 `binding:"required"`
	Kompensasi float64 `binding:"required"`
}
