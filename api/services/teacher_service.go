package services

import "github.com/kier1021/QuizMaker/api/repositories"

type TeacherService struct {
	teacherRepo *repositories.TeacherRepository
}

func NewTeacherService(teacherRepo *repositories.TeacherRepository) *TeacherService {
	return &TeacherService{
		teacherRepo: teacherRepo,
	}
}
