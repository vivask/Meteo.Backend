package dto

import "time"

type Pageable struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type Period struct {
	Begin time.Time `json:"begin"`
	End   time.Time `json:"end"`
}
