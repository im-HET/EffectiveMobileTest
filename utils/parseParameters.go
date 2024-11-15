package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func ParseParameters(r *http.Request, neededParams map[string]*Parameter) error {
	fmt.Println(r.URL)
	//Получаем query параметры
	queryParams, _ := url.ParseQuery(r.URL.RawQuery)

	//получаем body параметры
	jsonParams := make(map[string]interface{})
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return errors.New("Ошибка получения json " + err.Error())
	}
	if len(body) > 0 {
		err = json.Unmarshal(body, &jsonParams)
		if err != nil {
			return errors.New("Ошибка конвертации json " + err.Error())
		}
		fmt.Println(jsonParams)
	}

	for _, nParam := range neededParams { //Проверяем каждый ожидаемый параметр
		//Проверяем наличие параметра в path
		pathParam := r.PathValue(nParam.Name)
		if pathParam != "" {
			err := nParam.AddValue(pathParam)
			if err != nil {
				return errors.New("Ошибка добавления значения " + err.Error())
			}
			continue
		}

		if params, ok := queryParams[nParam.Name]; ok {
			for _, paramValue := range params {
				err := nParam.AddValue(paramValue)
				if err != nil {
					return errors.New("Ошибка добавления значения " + err.Error())
				}
			}
			continue
		}
		//Проверяем body параметры
		if jsonParam, ok := jsonParams[nParam.Name]; ok {
			err := nParam.AddValue(jsonParam)
			if err != nil {
				return errors.New("Ошибка добавления значения " + err.Error())
			}
			continue
		}

		//Проверяем обязательный ли параметр
		if nParam.Required {
			return errors.New("Недостаточно параметров")
		}
	}
	return nil
}
