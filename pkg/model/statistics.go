package model

type Statistics struct {
	Id          int    `json:"-" db:"id"`
	Status      string `json:"status" db:"status"`
	CreateCount int    `json:"create_count" db:"create_count"`
	UpdateCount int    `json:"update_count" db:"update_count"`
	DeleteCount int    `json:"delete_count" db:"delete_count"`
	ErrorCount  int    `json:"error_count" db:"error_count"`
}
