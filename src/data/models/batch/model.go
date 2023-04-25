package batch

import "time"

type Model struct {
	Idbatch     int64     `json:"IDBATCH"`
	Fecha_Batch time.Time `json:"FECHA_BATCH"`
	MontoTotal  float64   `json:"MONTO_TOTAL"`
	Editabatch  time.Time `json:"EDITABATCH"`
}

type Models []*Model
