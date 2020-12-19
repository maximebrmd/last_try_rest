package models

type Query struct {
	Filters map[string]interface{} `json:"filters"`
	Sort    map[string]interface{} `json:"sort"`
}
