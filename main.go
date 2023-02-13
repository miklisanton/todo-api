package main

import (
	"database/sql"
//	"log"
	"todo-api/controllers"

	//    "todo-api/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

func main() {
    r := httprouter.New()
    tc := controllers.NewTaskController(dbConnect())
    r.GET("/todos", tc.GetAll)
    r.GET("/todo/:id", tc.GetID)
    r.POST("/todos", tc.Create)
    r.DELETE("/todo/:id", tc.DeleteID)
    r.POST("/todo/:id", tc.UpdateID)
    http.ListenAndServe(":8080", r)
}

func dbConnect() *sql.DB {
    connStr := "user=zavropod dbname=todolist sslmode=disable password=123 port=5431"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err)
    }
    err = db.Ping()  
    if err != nil {
        panic(err)
    }
    return db
}
