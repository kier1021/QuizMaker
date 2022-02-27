package models

import "time"

type Teacher struct {
	ID           int       `json:"id"`
	FirstName    string    `json:"first_name"`
	MiddleName   string    `json:"middle_name"`
	LastName     string    `json:"last_name"`
	BirthDate    string    `json:"birth_date"`
	Address      string    `json:"address"`
	EmailAddress string    `json:"email_address"`
	PhoneNumber  string    `json:"phone_number"`
	Password     string    `json:"-"`
	CreatedDate  time.Time `json:"created_date"`
	UpdatedDate  time.Time `json:"updated_date"`
}

type TeacherSubjectBinding struct {
	TeacherID   int       `json:"teacher_id"`
	SubjectID   int       `json:"subject_id"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
}
