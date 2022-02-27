package models

import "time"

type Answer struct {
	ID          int       `json:"id"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
}

type AnswerItem struct {
	ID          int       `json:"id"`
	AnswerID    int       `json:"answer_id"`
	ItemID      int       `json:"item_id"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
}
