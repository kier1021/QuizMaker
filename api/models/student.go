package models

import "time"

type Student struct {
	ID           int       `json:"id"`
	FirstName    string    `json:"first_name"`
	MiddleName   string    `json:"middle_name"`
	LastName     string    `json:"last_name"`
	BirthDate    string    `json:"birth_date"`
	Address      string    `json:"address"`
	EmailAddress string    `json:"email_address"`
	PhoneNumber  string    `json:"phone_number"`
	SchoolName   string    `json:"school_name"`
	Password     string    `json:"-"`
	CreatedDate  time.Time `json:"created_date"`
	UpdatedDate  time.Time `json:"updated_date"`
}

type StudentSubjectBinding struct {
	StudentID   int       `json:"student_id"`
	SubjectID   int       `json:"subject_id"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
}

type StudentQuizBinding struct {
	StudentID   int       `json:"student_id"`
	QuizID      int       `json:"quiz_id"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
}

type StudentQuizAnswerBinding struct {
	StudentID   int       `json:"student_id"`
	QuizID      int       `json:"quiz_id"`
	AnswerID    int       `json:"answer_id"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
}
