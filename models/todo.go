package models

import (
    "fmt"
    "github.com/revel/revel"
)

type Todo struct {
    Name       string
    Completed  bool
    Number     int
    Creator    string
}

// "Implement" the Stringer interface: documentation of string verbs: https://golang.org/pkg/fmt/
func (t Todo) String() string {
    return fmt.Sprintf("ToDo: %s is done: %t", t.Name, t.Completed)
}

// Validation logic used within templates
func (todo *Todo) Validate(v *revel.Validation) {
    v.Required(todo.Name).Message("A ToDo has to have a name!")
}