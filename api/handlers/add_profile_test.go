package handlers

import (
	"students/gen/models"
	"sync"
	"testing"
)

var grade ="2nd"
var name = "hamza"
var class = "13"
var age int32 = 34
var student = &models.Students{Age: &age,
	ID:123,
	Grade:&grade,
	Class:&class,
	Name:&name}

var nilProfile = new(models.Students)

func TestAddProfileFunc(t *testing.T) {
	var profilesStore = make(map[int64]*models.Students)

	var profilesLock = &sync.Mutex{}
	t.Run("Add Profile", func(t *testing.T) {
		err := addStudentFun(student, profilesStore, profilesLock)
		if err != nil {
			t.Error("Test Failed: ", err)
			t.Fail()
		}
	})

	t.Run("Add Nil Profile", func(t *testing.T) {
		err := addStudentFun(nilProfile, profilesStore, profilesLock)
		if err == nil {
			t.Log("Test : Nil profile not added")
		} else {
			t.Log("Test : Nil profile added", err)
			t.Fail()
		}
	})

}