package controller

import (
	"html/template"
	"net/http"

	post "github.com/Tomoya185-miyawaki/go-blog/model"
)

func TopController(writer http.ResponseWriter, _ *http.Request) {
	posts := post.GetRows()
	t := template.Must(template.ParseFiles("template/top/index.html"))
	if err := t.Execute(writer, posts); err != nil {
		panic(err.Error())
	}
}
