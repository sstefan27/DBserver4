package rest

import (
	"encoding/json"
	"fmt"
	"go/problem4/db"
	"go/problem4/entity"
	"io/ioutil"
	"net/http"
)

func PostClass(rw http.ResponseWriter, r *http.Request) {
	reqBody := r.Body                         //request coming in body
	bodyBytes, err := ioutil.ReadAll(reqBody) //convert into bytes and check
	if hasError(rw, err, "Internal Issue") {
		return
	}
	var class entity.Class
	err = json.Unmarshal(bodyBytes, &class) //json->bytes
	if hasError(rw, err, "Internal Issue") {
		return
	}
	db.GetDB().Create(&class) //create
	fmt.Println(class)
	rw.Write(bodyBytes)
}

func GetClass(rw http.ResponseWriter, r *http.Request) {
	idClass := r.URL.Query().Get("id") //get the id from URL (link)
	var class entity.Class
	result := db.GetDB().Where("id= ?", idClass).Find(&class) //check if the result exists
	if result.RecordNotFound() {
		http.Error(rw, "Record not found", http.StatusInternalServerError)
		return
	}
	if result.Error != nil {
		http.Error(rw, "Internal issue", http.StatusInternalServerError)
		return
	}
	classData, _ := json.Marshal(class)
	rw.Write([]byte(classData))
}

func UpdateClass(rw http.ResponseWriter, r *http.Request) {
	reqBody := r.Body
	bodyBytes, err := ioutil.ReadAll(reqBody)
	if hasError(rw, err, "Internal Issue") {
		return
	}
	var class entity.Class
	err = json.Unmarshal(bodyBytes, &class)
	if hasError(rw, err, "Internal Issue") {
		return
	}
	var inputClass entity.Class
	json.Unmarshal(bodyBytes, &inputClass)
	db.GetDB().Model(&entity.Class{}).Updates(inputClass)

}
func DeleteClass(rw http.ResponseWriter, r *http.Request) {
	idClass := r.URL.Query().Get("id")
	result := db.GetDB().Where("id= ?", idClass).Delete(&entity.Class{})
	if result.Error != nil {
		http.Error(rw, "Inetrnal issue, please try again", http.StatusInternalServerError)
		return
	}
	rw.Write([]byte("Record deleted, succes"))
}
func GetClasses(rw http.ResponseWriter, r *http.Request) {
	var class []entity.Class
	db.GetDB().Find(&class)
	classData, _ := json.Marshal(class)
	rw.Write([]byte(classData))
}
