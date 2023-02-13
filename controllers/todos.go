package controllers

import (
	"database/sql"
	"net/http"
    "github.com/julienschmidt/httprouter"
    "log"
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
    if req.Method != http.MethodGet {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

    //create new
}
