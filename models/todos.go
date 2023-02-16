package models

import "todo-api/utils"

type Todo struct {
    Name string
    Description string
    Expires utils.CustomTime
    Priority int
}
