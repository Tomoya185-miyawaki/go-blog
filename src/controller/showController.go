package controller

import (
	"html/template"
	"net/http"

	post "github.com/Tomoya185-miyawaki/go-blog/model"
	path "github.com/Tomoya185-miyawaki/go-blog/utils"
)

func ShowController(writer http.ResponseWriter, req *http.Request) {
	id := path.GetShowPathId(req)
	post := post.GetRow(id)
	if !post.Valid {
		panic("レコードが存在しません")
	}
	t := template.Must(template.ParseFiles("template/show/index.html"))
	if err := t.Execute(writer, post.PostDisplay); err != nil {
		panic(err.Error())
	}
}
