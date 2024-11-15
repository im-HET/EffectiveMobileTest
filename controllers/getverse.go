package controllers

import (
	"mediaLibrary_v2/services"
	"mediaLibrary_v2/utils"
	"net/http"
)

func GetVerse(w http.ResponseWriter, r *http.Request) {
	i := services.NewGetVerse()
	err := utils.ParseParameters(
		r,                         //Передаем request
		i.GetExpectedParameters(), //и слайс с параметрами
		//если параметр присутствует в строке запроса он получит значение
		//или массив значений
	)
	if err != nil { //Если на этапе обработки url и получения значений произошла ошибка
		http.Error(w, err.Error(), 400) //то отправляем клиенту ошибку 400
		return
	}
	result, err := i.Get()
	if err != nil { //Если в работе сервиса произошла ошибка
		http.Error(w, err.Error(), 500) //то отправляем клиенту ошибку 500
		return
	}
	w.Write(result)
}
