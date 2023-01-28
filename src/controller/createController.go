package controller

import (
	"html/template"
	"net/http"
)

func CreateController(writer http.ResponseWriter, _ *http.Request) {
	t, err := template.ParseFiles("template/create/index.html")
	if err != nil {
		panic(err.Error())
	}
	if err := t.Execute(writer, nil); err != nil {
		panic(err.Error())
	}
}
