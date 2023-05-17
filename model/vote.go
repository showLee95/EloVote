package model

type Vote struct {
	VId   int64   `json:"vid" db:"vid"`
	Name  string  `json:"votename" db:"votename"`
	Vote  int64   `json:"vote" db:"vote"`
	Rate  float64 `json:"rate" db:"rate"`
	Score int64   `json:"score" db:"score"`
}

type Votename struct {
	Victory string `json:"victory"`
	Burden  string `json:"burden"`
}
