package controllers

import (
    "github.com/revel/revel"
    "todo/models"
    "todo/app/routes"
)

type App struct {
	*revel.Controller
}

// This Action function renders the .../view/App/index.html
func (c App) Index() revel.Result {
	return c.Render()
}

// This Action function handles the login of an user
func (c App) Login(user *models.User) revel.Result {
    user.Validate(c.Validation)

    // Handle errors
    if c.Validation.HasErrors() {
        c.Validation.Keep()
        c.FlashParams()
        return c.Redirect(App.Index)
    }

    // Storing the user's name in the session
    c.Session["userName"] = user.Name
    c.Session.SetNoExpiration()

    // Inform the user about the success on the next page
    c.Flash.Success("Welcome, " + user.Name + "!")

    return c.Redirect(Todo.Index)
}

// This Action function is checking whether or not a user have already been logged in
func (c App) CheckLogin() revel.Result {
    if _, ok := c.Session["userName"]; ok {
        c.Flash.Success("You are already logged in " + c.Session["userName"] + "!")
        return c.Redirect(routes.Todo.Index())
    }

    return nil
}