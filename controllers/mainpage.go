package controllers

import (
	"html/template"
	"net/http"
)

type viewData struct {
	Title string
}

func Mainpage(w http.ResponseWriter, r *http.Request) {
	data := viewData{
		Title: "Онлайн библиотека Effective Mobile",
	}
	tmpl, _ := template.ParseFiles("html/index.html")
	tmpl.Execute(w, data)
}
