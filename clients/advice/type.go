package advice

type Slip struct {
	Advice Advice `json:"slip"`
}

type Advice struct {
	ID     int    `json:"id"`
	Advice string `json:"advice"`
}
