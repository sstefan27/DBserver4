package rest

import (
	"encoding/json"
	"fmt"
	"go/problem4/db"
	"go/problem4/entity"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

func PostStudent(rw http.ResponseWriter, r *http.Request) {
	reqBody := r.Body                         //request coming in body
	bodyBytes, err := ioutil.ReadAll(reqBody) //convert into bytes and check
	if hasError(rw, err, "Internal Issue") {
		return
	}
	var student entity.Student
	err = json.Unmarshal(bodyBytes, &student) //json->bytes
	if hasError(rw, err, "Internal Issue") {
		return
	}
	db.GetDB().Create(&student) //create
	fmt.Println(student)
	rw.Write(bodyBytes)
}

func GetStudent(rw http.ResponseWriter, r *http.Request) {
	idStudent := r.URL.Query().Get("id") //get the id from URL (link)
	var student entity.Student
	result := db.GetDB().Where("id=?", idStudent).Find(&student) //check if the result exists
	if result.RecordNotFound() {
		http.Error(rw, "Record not found", http.StatusInternalServerError)
		return
	}
	if result.Error != nil {
		http.Error(rw, "Internal issue", http.StatusInternalServerError)
		return
	}
	studentData, _ := json.Marshal(student)
	rw.Write([]byte(studentData))
}

func UpdateStudent(rw http.ResponseWriter, r *http.Request) {
	reqBody := r.Body
	bodyBytes, err := ioutil.ReadAll(reqBody)
	if hasError(rw, err, "Internal Issue") {
		return
	}
	var student entity.Student
	err = json.Unmarshal(bodyBytes, &student)
	if hasError(rw, err, "Internal Issue") {
		return
	}
	var inputStudent entity.Student
	json.Unmarshal(bodyBytes, &inputStudent)
	db.GetDB().Model(&entity.Student{}).Updates(inputStudent)

}
func DeleteStudent(rw http.ResponseWriter, r *http.Request) {
	idStudent := r.URL.Query().Get("id")
	result := db.GetDB().Where("id= ?", idStudent).Delete(&entity.Student{})
	if result.Error != nil {
		http.Error(rw, "Inetrnal issue, please try again", http.StatusInternalServerError)
		return
	}
	rw.Write([]byte("Record deleted, succes"))
}

func GetStudents(rw http.ResponseWriter, r *http.Request) {
	var student []entity.Student
	db.GetDB().Find(&student)
	studentData, _ := json.Marshal(student)
	rw.Write([]byte(studentData))
}

func EnrollStudent(rw http.ResponseWriter, r *http.Request) {
	reqBody := r.Body
	bodyBytes, err := ioutil.ReadAll(reqBody)
	if hasError(rw, err, "Internal Issue") {
		return
	}
	var student entity.Student
	err = json.Unmarshal(bodyBytes, &student)
	if hasError(rw, err, "Internal Issue") {
		return
	}
	//db.GetDB().Preload("class_students").Find(&student)
	db.GetDB().Model(student).Association("Class").Replace(student.Class)
	rw.Write([]byte("Enrolled succesfully done"))
}
func hasError(rw http.ResponseWriter, err error, message string) bool {
	logger := new(logrus.Entry)
	if err != nil {
		logger.WithError(err).Error(message)
		rw.Write([]byte(message))
		return true
	}
	return false
}
