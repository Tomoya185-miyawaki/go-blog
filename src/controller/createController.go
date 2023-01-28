package controller

import (
	"html/template"
	"log"
	"net/http"

	post "github.com/Tomoya185-miyawaki/go-blog/model"
)

func CreateController(writer http.ResponseWriter, req *http.Request) {
	// GETメソッドの場合
	if req.Method == http.MethodGet {
		t, err := template.ParseFiles("template/create/index.html")
		if err != nil {
			panic(err.Error())
		}
		if err := t.Execute(writer, nil); err != nil {
			panic(err.Error())
		}
		return
	}

	// POSTメソッドの場合
	if req.Method == http.MethodPost {
		post.Create(req.FormValue("title"), req.FormValue("content"))
		http.Redirect(writer, req, "/", 302)
		return
	}

	log.Println(http.MethodGet)
	panic("対応していないHTTPメソッドです")
}
