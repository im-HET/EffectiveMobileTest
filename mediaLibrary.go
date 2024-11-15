package main

import (
	"fmt"
	"mediaLibrary_v2/db"
	"net/http"

	"mediaLibrary_v2/controllers"
	"mediaLibrary_v2/settings"

	//"github.com/go-chi/chi"
	_ "github.com/lib/pq"
	//_ "github.com/swaggo/http-swagger/example/go-chi/docs"
	"github.com/swaggo/http-swagger"
)

// @title Тестовое задание "Медиа библиотека"
// @version 2
// @description Тестовое задание в EffectiveMobile "Медиа библиотека"
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url t.me/@Gmsviridov
// @contact.email gmsviridov@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host http://localhost:8080
// @BasePath /v2
func main() {
	settings.Load()
	db.StartDB()
	fmt.Println("start service")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /favicon.ico", controllers.Favicon)
	mux.HandleFunc("GET /", controllers.Mainpage)
	mux.HandleFunc("GET /items", controllers.GetItems)
	mux.HandleFunc("GET /items/{id}", controllers.GetItem)
	mux.HandleFunc("GET /items/{id}/text/{num}", controllers.GetVerse)
	mux.HandleFunc("PUT /items", controllers.PutItem)
	mux.HandleFunc("DELETE /items/{id}", controllers.DelItem)
	mux.HandleFunc("PATCH /items/{id}", controllers.PatchItem)
	mux.HandleFunc("GET /info", controllers.GetInfo)
	mux.HandleFunc("GET /swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/docs/swagger.json")))
	http.ListenAndServe(settings.Port, mux)
	//fmt.Println("web server запущен")
}
