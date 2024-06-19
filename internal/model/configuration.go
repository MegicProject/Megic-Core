package model

type ConfigurationModel struct {
	Code  string `db:"code" json:"code"`
	Value string `db:"value" json:"value"`
}

type HistoryPart struct {
	Text string `json:"text"`
}

type HistoryItem struct {
	Parts []HistoryPart `json:"parts"`
	Role  string        `json:"role"`
}
