package models

import (
    "github.com/revel/revel"
    "fmt"
)

type User struct {
    Name string
}

// "Implement" the Stringer interface: documentation of string verbs: https://golang.org/pkg/fmt/
func (u *User) String() string {
    return fmt.Sprintf("User(%s)", u.Name)
}

// Validation logic used within templates
func (user *User) Validate(v *revel.Validation) {
    v.Required(user.Name).Message("A User Name is required!")
    v.MinSize(user.Name, 4).Message("The User Name must be at least 4 characters long!")
    v.MaxSize(user.Name, 10).Message("The User Name may not be longer than 10 characters!")
}