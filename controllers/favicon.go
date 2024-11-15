package controllers

import (
	"net/http"
)

func Favicon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/favicon.ico")
}
