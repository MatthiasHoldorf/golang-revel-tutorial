package controllers

import (
    "github.com/revel/revel"
    "todo/models"
    "todo/app/routes"
    "strconv"
)

type Todo struct {
    *revel.Controller
}

// The list to store added tasks
var todoList = []models.Todo{}

// This Action function renders the .../view/Todo/index.html
func (c Todo) Index() revel.Result {
    return c.Render(todoList)
}

// This Action function is adding a task to the todoList declared above
func (c Todo) AddTodo(todo models.Todo) revel.Result {
    todo.Validate(c.Validation)

    // Handle errors
    if c.Validation.HasErrors() {
        c.Validation.Keep()
        c.FlashParams()
        return c.Redirect(Todo.Index)
    }

    // Setting the Number and Creator field of the task accordingly
    todo.Number = len(todoList)
    todo.Number++
    todo.Creator = c.Session["userName"]
    todoList = append(todoList, todo)

    return c.Redirect(Todo.Index)
}

// This Action function completes one or more tasks from the todolist
func (c Todo) CompleteToDo() revel.Result {
    c.Request.ParseForm()
    for i, _ := range c.Request.Form {
        t, _ := strconv.Atoi(i)
        t--
        todoList[t].Completed = true
    }

    return c.Redirect(Todo.Index)
}

// This Action function is checking whether or not a user have already been logged in
func (c Todo) CheckLogin() revel.Result {
    if _, ok := c.Session["userName"]; ok {
        return nil
    }

    c.Flash.Error("Please log in first!")
    return c.Redirect(routes.App.Index())
}

// This Action function logs out a user and destroys the session associated
func (c Todo) Logout() revel.Result {
    for k := range c.Session {
        delete(c.Session, k)
    }
    return c.Redirect(App.Index)
}