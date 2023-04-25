package switchboardItems

type Model struct {
	SwitchboardID int64  `json:"SwitchboardID"`
	ItemNumber    int64  `json:"ItemNumber"`
	ItemText      string `json:"ItemText"`
	Command       int64  `json:"Command"`
	Argument      string `json:"Argument"`
}

type Models []*Model
