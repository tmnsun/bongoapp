package app

import (
	"appengine"
	"appengine/datastore"
	"appengine/user"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"strconv"
)

type Page struct {
	Title string
}

type Task struct {
	Id           int64 `datastore:"-"` // instructs datastore to ignore
	Title        string
	Details		 string
	Category     string
	State		 string
	Dt_completed int64
	Dt_created   int64
}

var cached_templates = template.Must(template.ParseGlob("app/templates/*.html"))

func init() {
	http.HandleFunc("/api/", router)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/update", update)
	http.HandleFunc("/", home)
}

func router(res http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case "GET":
		get(res, req)
	case "POST":
		post(res, req)
	case "PUT":
		put(res, req)
	case "DELETE":
		archive(res, req)
	default:
		fmt.Fprintf(res, "{}")
	}
}

func home(res http.ResponseWriter, req *http.Request) {
	page := Page{}
	page.Title = "bongo app"
	renderTemplate(res, "index.html", &page)
}

func logout(res http.ResponseWriter, req *http.Request) {
	c := appengine.NewContext(req)
	logout_url, _ := user.LogoutURL(c, "/")
	http.Redirect(res, req, logout_url, http.StatusTemporaryRedirect)
}

func get(res http.ResponseWriter, req *http.Request) {
	c := appengine.NewContext(req)

	q := datastore.NewQuery("Task").Filter("State =", "active").Order("Title").Limit(50)
	tasks := make([]Task, 0, 50)
	keys, err := q.GetAll(c, &tasks)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	for i, key := range keys {
		tasks[i].Id = key.IntID()
	}

	data, _ := json.Marshal(tasks)
	fmt.Fprintf(res, string(data))
}

func post(res http.ResponseWriter, req *http.Request) {
	c := appengine.NewContext(req)
	var task Task
	var model = req.FormValue("model")
	json.Unmarshal([]byte(model), &task)

	// write to data store
	_, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Task", nil), &task)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	// return Ok
	fmt.Fprintf(res, "{'Ok'}")
}

func put(res http.ResponseWriter, req *http.Request) {
	c := appengine.NewContext(req)
	var task Task
	var model = req.FormValue("model")
	json.Unmarshal([]byte(model), &task)

	key := datastore.NewKey(c, "Task", "", task.Id, nil)
	_, err := datastore.Put(c, key, &task)
	if err != nil {
		c.Errorf(err.Error())
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(res, "'{'Ok'}")
}


func archive(res http.ResponseWriter, req *http.Request) {

	c := appengine.NewContext(req)
	taskId, _ := strconv.ParseInt(strings.Replace(req.URL.Path, "/api/", "", 1), 10, 64)

	key := datastore.NewKey(c, "Task", "", taskId, nil)
	task := new(Task)
	err := datastore.Get(c, key, task)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	task.State = "archived"
	_, err2 := datastore.Put(c, key, task)
	if err2 != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(res, "'{'Ok'}")
}

func update(res http.ResponseWriter, req *http.Request) {

	c := appengine.NewContext(req)
	q := datastore.NewQuery("Task").Limit(50)
	tasks := make([]Task, 0, 50)

	keys, err := q.GetAll(c, &tasks)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	for i, key := range keys {
		if tasks[i].State == "" {
			tasks[i].State = "active"
		}

		_, err := datastore.Put(c, key, &tasks[i])
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
	}
	fmt.Fprintf(res, "'{'Ok'}")
}

func renderTemplate(res http.ResponseWriter, template string, p *Page) {

	err := cached_templates.ExecuteTemplate(res, template, p)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}
