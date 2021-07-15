package db

import "go/problem4/entity"

func InitData() {
	student := &entity.Student{
		StudentID: "2",
		LastName:  "DummyLast",
		FirstName: "FirstDummy",
		Age:       42,
	}
	GetDB().Save(&student)
}
