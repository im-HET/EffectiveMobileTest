package services

import (
	"encoding/json"
	"errors"
	"mediaLibrary_v2/db"

	"mediaLibrary_v2/utils"
	"strings"
)

type GetVerse struct {
	service
}

func NewGetVerse() GetVerse {
	i := GetVerse{}
	id := utils.Parameter{
		Name:      "id",
		Required:  true,
		Condition: " = ",
		Column:    "library.id",
		Validator: utils.ValidateInt,
	}
	num := utils.Parameter{
		Name:      "num",
		Required:  true,
		Condition: " = ",
		Column:    "library.text",
		Validator: utils.ValidateInt,
	}
	i.expectedParameters = map[string]*utils.Parameter{
		id.Name:  &id,
		num.Name: &num,
	}
	return i
}

func (i GetVerse) Get() ([]byte, error) {
	type responseBeforeJson struct {
		Amount *int    `db:"amount"`
		Num    *int    `db:"num"`
		Verse  *string `db:"verse"`
	}
	sql := []string{
		"SELECT *",
		"FROM getverse(",
		i.expectedParameters["id"].GetValue(),
		", ",
		i.expectedParameters["num"].GetValue(),
		")",
	}
	resp := responseBeforeJson{}
	err := db.SqlGet(strings.Join(sql, " "), &resp)
	if err != nil {
		return nil, err
	}
	resultJson, err := json.Marshal(resp)
	if err != nil {
		return nil, errors.New("ошибка серилизации в json " + err.Error())
	}
	return resultJson, nil
}
