package establishment

type Model struct {
	Establishmentid int64       `json:"ESTABLISHMENTID"`
	NomEsta         string      `json:"NOM_ESTA"`
	NomDueno        string      `json:"NOM_DUENO"`
	Telefono        string      `json:"TELEFONO"`
	CalAmbiental    string      `json:"CAL_AMBIENTAL"`
	SsPatronal      string      `json:"SS_PATRONAL"`
	EAddress1       string      `json:"E_ADDRESS#1"`
	ETownid1        string      `json:"E_TOWNID1"`
	EZipCode1       interface{} `json:"E_ZIP_CODE1"`
	EAddress2       string      `json:"E_ADDRESS#2"`
	ETownid2        string      `json:"E_TOWNID2"`
	EZipCode2       string      `json:"E_ZIP_CODE2"`
	Notas           interface{} `json:"NOTAS"`
	Chofer          bool        `json:"CHOFER"`
	Tipo            interface{} `json:"TIPO"`
	Precio          interface{} `json:"PRECIO"`
	ECountry1       string      `json:"E_COUNTRY1"`
	ECountry2       string      `json:"E_COUNTRY2"`
	ChkRuben        bool        `json:"Chk_Ruben"`
	Tanques         string      `json:"Tanques"`
	Drone           interface{} `json:"Drone"`
	Reembolso       interface{} `json:"Reembolso"`
	Reembolso1      interface{} `json:"Reembolso1"`
	NumVendor       interface{} `json:"Num_Vendor"`
	Carta           interface{} `json:"CARTA"`
	Filtros         interface{} `json:"FILTROS"`
}

type Models []*Model
