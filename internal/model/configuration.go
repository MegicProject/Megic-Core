package model

type ConfigurationModel struct {
	Code  string `db:"CODE" json:"code"`
	Value string `db:"VALUE" json:"value"`
}

type HistoryPart struct {
	Text string `json:"text"`
}

type HistoryItem struct {
	Parts []HistoryPart `json:"parts"`
	Role  string        `json:"role"`
}
