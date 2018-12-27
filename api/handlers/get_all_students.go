package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"students/gen/models"
	"students/gen/restapi/operations/students"
)

func NewGetAllStudents(database map[int64]*models.Students) students.GetAllStudentsHandler{
	return &getAllStudent{database:database}
}

type getAllStudent struct {
	database map[int64]*models.Students
}

func (fn *getAllStudent) Handle(params students.GetAllStudentsParams) middleware.Responder {
	return students.NewGetAllStudentsOK().WithPayload(GetAllStudentFun(fn.database))
}

func GetAllStudentFun(database map[int64]*models.Students) (result []*models.Students){
	for _,student := range database {
		result = append(result,student)
	}
	return
}
