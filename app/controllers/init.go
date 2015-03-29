package controllers

import (
"github.com/revel/revel"
)

// Registration of interceptor methods
func init() {
    revel.InterceptMethod(Todo.CheckLogin, revel.BEFORE)
}
