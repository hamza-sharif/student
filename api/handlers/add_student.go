package handlers

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"math/rand"
	"students/gen/models"
	"students/gen/restapi/operations/students"
	"sync"
)

func NewAddStudent(database map[int64]*models.Students, lock *sync.Mutex) students.AddStudentHandler {
	return &addStudent{database: database,
		lock: lock}
}

type addStudent struct {
	database map[int64]*models.Students
	lock     *sync.Mutex
}

func (a *addStudent) Handle(params students.AddStudentParams) middleware.Responder {
	if err := addStudentFun(params.Body, a.database, a.lock); err != nil {
		return students.NewAddStudentDefault(500).
			WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
	}
	return students.NewAddStudentCreated().WithPayload(params.Body)
}

func addStudentFun(student *models.Students, database map[int64]*models.Students,
	lock *sync.Mutex) error {
	if database == nil {
		return errors.New(500, "profileStore must be present")
	}
	lock.Lock()
	defer lock.Unlock()
	id := rand.Int63()
	student.ID = id
	database[id] = student
	return nil
}
