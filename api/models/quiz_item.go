package models

import "time"

type MultipleChoiceItem struct {
	ID          int       `json:"id"`
	Question    string    `json:"question"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
}

type MultipleChoiceAnswer struct {
	ID          int       `json:"id"`
	ItemID      string    `json:"item_id"`
	Answer      string    `json:"answer"`
	Label       string    `json:"label"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
}

type IdentificationItem struct {
	ID          int       `json:"id"`
	Question    string    `json:"question"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
}

type IdentificationAnswer struct {
	ID          int       `json:"id"`
	ItemID      string    `json:"item_id"`
	Answer      string    `json:"answer"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
}

type FillInBlankItem struct {
	ID          int       `json:"id"`
	Question    string    `json:"question"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
}

type FillInBlankAnswer struct {
	ID          int       `json:"id"`
	ItemID      string    `json:"item_id"`
	Answer      string    `json:"answer"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
}
