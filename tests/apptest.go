package tests

import (
    "github.com/revel/revel"
    "net/url"
)

type AppTest struct {
	revel.TestSuite
}

func (t *AppTest) Before() {

	println("Set up")
}


func (t *AppTest) TestThatAppIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t *AppTest) TestThatAddTodoWorks() {
    t.PostForm("/Todo/AddTodo", url.Values{
        "todo.Name" : {"Read Emails"},
    })
    t.AssertStatus(201)
}

func (t *AppTest) After() {
	println("Tear down")
}
