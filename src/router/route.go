package router

import (
	"net/http"

	controller "github.com/Tomoya185-miyawaki/go-blog/controller"
)

func Route() {
	http.HandleFunc("/", controller.TopController)
	http.HandleFunc("/create", controller.CreateController)
	http.ListenAndServe(":8000", nil)
}
