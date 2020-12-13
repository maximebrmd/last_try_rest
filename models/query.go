package models

type Query struct {
	Filters map[string]interface{} `json:"filters" binding:"required"`
	Sort    map[string]interface{} `json:"sort" binding:"required"`
}
