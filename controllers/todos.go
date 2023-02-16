package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"todo-api/models"
    "github.com/lib/pq"
	"github.com/julienschmidt/httprouter"
)

type TaskController struct {
    db *sql.DB
}

func NewTaskController(db *sql.DB) *TaskController {
    return &TaskController{db}
}

func (tc TaskController) GetAll(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
    if req.Method != http.MethodGet {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

    //getting all functionality
}

func (tc TaskController) GetID(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
    if req.Method != http.MethodGet {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

    //Get by id 
    id := ps.ByName("id")
    log.Println(id)
}

func (tc TaskController) UpdateID(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
    if req.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

    //Update by id
    id := ps.ByName("id")
    log.Println(id)
}

func (tc TaskController) DeleteID(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
    if req.Method != http.MethodGet {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

    //delete by id
    id := ps.ByName("id")
    log.Println(id)
}

func (tc TaskController) Create(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
    if req.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

    var tasks []models.Todo
    d := json.NewDecoder(req.Body) 
    //Populate data into slice
    err := d.Decode(&tasks)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        log.Println(err)
        return
    }
    //Prepare statement
    txn, err := tc.db.Begin()
    if err != nil {
        w.WriteHeader(500)
        log.Println(err)
        return
    }
    stmt, err := txn.Prepare(pq.CopyIn("tasks", "name", "description", "expires", "priority"))
    if err != nil {
        w.WriteHeader(500)
        log.Println(err)
        return
    }
    //Check data and insert rows 
    for _, t := range tasks {
        if (t.Name == "" || t.Priority > 3 || t.Priority < 0) {
           w.WriteHeader(http.StatusBadRequest)
           log.Printf("Wrong format %v", t)
           return
       }
       //execute statement
       _, err = stmt.Exec(t.Name, t.Description, t.Expires, t.Priority) 
       if err != nil {
           w.WriteHeader(500)
           log.Println(err)
       }
   }
   //Flush buffered data
   _, err = stmt.Exec()
   if err != nil {
       log.Fatal(err)
   }
   
   err = stmt.Close()
   if err != nil {
       log.Fatal(err)
   }
    
   err = txn.Commit()
   if err != nil {
       log.Fatal(err)
   }
   //Send statusOK
   w.WriteHeader(200)
   return
}
