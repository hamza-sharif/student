package main

import (
	"flag"
	"github.com/go-openapi/loads"
	"log"
	"students/gen/models"
	"students/gen/restapi"
	"students/gen/restapi/operations"
	"students/api/handlers"
	"sync"
)

var portflag = flag.Int("port", 3000, "Port for server")
var studentStore = make(map[int64]*models.Students)
var lock = &sync.Mutex{}

func main() {

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}
	api := operations.NewStudentInfoAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer func() {
		_ = server.Shutdown()
	}()

	flag.Parse()
	server.EnabledListeners = []string{"http"}
	server.Port = *portflag

	api.StudentsAddStudentHandler = handlers.NewAddStudent(studentStore, lock)
	api.StudentsGetAllStudentsHandler = handlers.NewGetAllStudents(studentStore)

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

	}
