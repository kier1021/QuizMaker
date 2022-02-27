package controllers

import "github.com/kier1021/QuizMaker/api/services"

type TeacherController struct {
	teacherSrv *services.TeacherService
}

func NewTeacherController(teacherSrv *services.TeacherService) *TeacherController {
	return &TeacherController{
		teacherSrv: teacherSrv,
	}
}
