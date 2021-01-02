package model

type Statistics struct {
	CreateCount int `json:"create_count"`
	UpdateCount int `json:"update_count"`
	DeleteCount int `json:"delete_count"`
	ErrorCount  int `json:"error_count"`
}
