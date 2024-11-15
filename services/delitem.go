package services

import (
	"encoding/json"
	"errors"
	"mediaLibrary_v2/db"

	"mediaLibrary_v2/utils"
	"strings"
)

type DelItem struct {
	service
}

func NewDelItem() DelItem {
	i := DelItem{}
	id := utils.Parameter{
		Name:      "id",
		Required:  true,
		Condition: " = ",
		Column:    "library.id",
		Validator: utils.ValidateInt,
	}
	i.expectedParameters = map[string]*utils.Parameter{id.Name: &id}
	return i
}

func (d DelItem) Del() ([]byte, error) {
	type responseBeforeJson struct { //используем указатели для того
		Id    *int `db:"id"` //что бы без ошибок обрабатывать null значения из базы
		Group *int `db:"groupid"`
	}
	resp := responseBeforeJson{}
	sql := []string{
		"SELECT",
		"id,",
		"groupid",
		"FROM library",
		utils.GetWhere(
			"",
			*d.expectedParameters["id"],
		)}
	err := db.SqlGet(strings.Join(sql, " "), &resp)
	if err != nil {
		return nil, err
	}
	sql = []string{
		"DELETE FROM",
		"library",
		utils.GetWhere(
			"",
			*d.expectedParameters["id"],
		),
		"RETURNING id",
	}
	err = db.SqlGet(strings.Join(sql, " "), &resp.Id)
	if err != nil {
		return nil, err
	}
	resultJson, err := json.Marshal(resp)
	if err != nil {
		return nil, errors.New("ошибка серилизации в json " + err.Error())
	}
	return resultJson, nil
}
