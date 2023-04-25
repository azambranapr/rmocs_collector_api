package transaccion

type Model struct {
	Transaccionid   int64       `json:"TRANSACCIONID"`
	Telefono        string      `json:"TELEFONO"`
	NumFactura      int64       `json:"NUM_FACTURA"`
	Costogalones    float64     `json:"COSTOGALONES"`
	Galones         int64       `json:"GALONES"`
	Costototal      int64       `json:"COSTOTOTAL"`
	FactPagada      bool        `json:"FACT_PAGADA"`
	Fecha           string      `json:"FECHA"`
	ClasePago       bool        `json:"CLASE_PAGO"`
	NumCheque       int64       `json:"NUM_CHEQUE"`
	CasoEspecial    bool        `json:"CASO_ESPECIAL"`
	Hacienda        bool        `json:"Hacienda"`
	Certi           interface{} `json:"Certi"`
	FechaReembolso  interface{} `json:"FechaReembolso"`
	CostPaid        interface{} `json:"CostPaid"`
	TotalPaid       int64       `json:"TotalPaid"`
	NumChk          int64       `json:"NumChk"`
	FechaReembolso1 interface{} `json:"FechaReembolso1"`
	CostPaid1       interface{} `json:"CostPaid1"`
	TotalPaid1      int64       `json:"TotalPaid1"`
	NumChk1         int64       `json:"NumChk1"`
	Tipoalma        string      `json:"TIPOALMA"`
	IDBatch         interface{} `json:"IDBatch"`
	EditBatch       string      `json:"EditBatch"`
	OleinCK         bool        `json:"OleinCK"`
}

type Models []*Model
