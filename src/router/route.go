package router

import (
	"net/http"

	controller "github.com/Tomoya185-miyawaki/go-blog/controller"
)

func Route() {
	http.HandleFunc("/show/", controller.ShowController)
	http.HandleFunc("/create", controller.CreateController)
	http.HandleFunc("/", controller.TopController)
	http.ListenAndServe(":8000", nil)
}
