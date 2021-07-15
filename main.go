package main

import (
	"go/problem4/config"
	"go/problem4/db"
	"go/problem4/rest"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Not able to load env file")
	}
	err = config.InitConfig()
	if err != nil {
		log.Fatal("Not able to create config")
	}
	db.InitDatabase()
	db.InitData()
	router := chi.NewRouter()
	router.Route("/"+config.GetConfig().APPVersion, func(r chi.Router) {
		//student
		r.Get("/student", rest.GetStudent)
		r.Post("/student", rest.PostStudent)
		r.Put("/student", rest.UpdateStudent)
		r.Delete("/student", rest.DeleteStudent)
		r.Get("/students", rest.GetStudents)

		//class

		r.Get("/class", rest.GetClass)
		r.Post("/class", rest.PostClass)
		r.Put("/class", rest.UpdateClass)
		r.Delete("/class", rest.DeleteClass)
		r.Get("/classes", rest.GetClasses)
		r.Put("/enroll", rest.EnrollStudent)

		r.Post("/login", rest.Login)
		r.Get("/welcome", rest.Welcome)
		r.Get("/logout", rest.Logout)

	})
	http.ListenAndServe(":"+config.GetConfig().Port, router)
}
