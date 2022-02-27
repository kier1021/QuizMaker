package models

import "time"

type Quiz struct {
	ID          int           `json:"id"`
	SubjectID   int           `json:"subject_id"`
	Name        string        `json:"name"`
	Level       string        `json:"level"`
	TimeLimit   time.Duration `json:"time_limit"`
	Deadline    time.Time     `json:"deadline"`
	CreatedDate time.Time     `json:"created_date"`
	UpdatedDate time.Time     `json:"updated_date"`
}
