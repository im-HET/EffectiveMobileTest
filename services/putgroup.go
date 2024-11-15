package services

import (
	"encoding/json"
	"errors"
	"mediaLibrary_v2/db"
	"mediaLibrary_v2/utils"
	"strings"
)

type PutGroup struct {
	service
}

func NewPutGroup() PutGroup {
	p := PutGroup{}
	group := utils.Parameter{
		Name:      "group",
		Required:  true,
		Condition: " = ",
		Column:    "name",
		Validator: utils.ValidateString,
	}
	p.expectedParameters = map[string]*utils.Parameter{
		group.Name: &group,
	}
	return p
}

func (p PutGroup) Put() ([]byte, error) {
	type responseBeforeJson struct {
		Id *int `db:"id"`
	}
	//получаем Id группы
	var groupId int = 0
	var err error
	sql := []string{
		"SELECT",
		"groups.id",
		"FROM groups",
		utils.GetWhere(
			"",
			*p.expectedParameters["group"],
		),
	}
	err = db.SqlGet(strings.Join(sql, " "), &groupId)
	if err != nil {
		if errors.Is(err, db.ErrNoRows{}) { //если группа не была найдена по имени то добавляем новую
			sql = []string{
				"INSERT INTO",
				"groups",
				utils.GetColumns(p.expectedParameters),
				"VALUES",
				utils.GetValues(p.expectedParameters),
				"RETURNING id",
			}
			err = db.SqlExec(strings.Join(sql, " "), &groupId)
			if err != nil {
				return nil, errors.New("ошибка добавления группы " + err.Error())
			}
		} else {
			return nil, errors.New("ошибка получения id группы " + err.Error())
		}
	}
	resp := responseBeforeJson{Id: &groupId}
	resultJson, err := json.Marshal(resp)
	if err != nil {
		return nil, errors.New("ошибка серилизации в json " + err.Error())
	}
	return resultJson, nil
}
