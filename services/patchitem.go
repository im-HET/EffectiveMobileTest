package services

import (
	"encoding/json"
	"errors"
	"mediaLibrary_v2/db"
	"mediaLibrary_v2/utils"
	"strings"
)

type PatchItem struct {
	service
}

func NewPatchItem() PatchItem { //Конструктор для сервиса выдачи списка песен
	i := PatchItem{}
	song := utils.Parameter{ //Определяем ожидаемые/возможные параметры для данного сервиса
		Name:      "song",
		Required:  false, //Если наличие параметра обязательно для работы сервиса ставим true
		Condition: " = ",
		Validator: utils.ValidateString, //Валидатор параметра. можно добавить свой или выбрать "стандартный"
		Column:    "song",
	}
	group := utils.Parameter{
		Name:      "group",
		Required:  false,
		Condition: " = ",
		Column:    "groupid",
		Validator: utils.ValidateString,
	}
	id := utils.Parameter{
		Name:      "id",
		Required:  true,
		Condition: " = ",
		Column:    "id",
		Validator: utils.ValidateInt,
	}
	releseDate := utils.Parameter{
		Name:      "releasedate",
		Required:  false,
		Condition: " = ",
		Column:    "releasedate",
		Validator: utils.ValidateDate,
	}
	link := utils.Parameter{
		Name:      "link",
		Required:  false,
		Condition: " = ",
		Column:    "link",
		Validator: utils.ValidateString,
	}
	text := utils.Parameter{
		Name:      "text",
		Required:  false,
		Condition: " = ",
		Column:    "text",
		Validator: utils.ValidateString,
	}

	i.expectedParameters = map[string]*utils.Parameter{
		id.Name:         &id,
		song.Name:       &song,
		group.Name:      &group,
		releseDate.Name: &releseDate,
		link.Name:       &link,
		text.Name:       &text,
	}

	return i
}

func (p PatchItem) Patch() ([]byte, error) {
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
			*p.expectedParameters["id"],
		)}
	err := db.SqlGet(strings.Join(sql, " "), &resp)
	if err != nil {
		return nil, err
	}

	if p.expectedParameters["group"].GetValue() != "" {
		pgroup := NewPutGroup()
		pg := pgroup.expectedParameters["group"]
		pg.AddValue(p.expectedParameters["group"].GetValue())
		b, err := pgroup.Put()
		if err != nil {
			return nil, err
		}
		json.Unmarshal(b, &resp.Group)
		p.expectedParameters["group"].Clear()
		p.expectedParameters["group"].AddValue(*resp.Group)
	}
	sql = []string{
		"UPDATE library",
		utils.GetSet(
			*p.expectedParameters["song"],
			*p.expectedParameters["group"],
			*p.expectedParameters["releasedate"],
			*p.expectedParameters["link"],
			*p.expectedParameters["text"],
		),
		utils.GetWhere(
			"",
			*p.expectedParameters["id"],
		),
		"RETURNING id",
	}
	err = db.SqlExec(strings.Join(sql, " "), &resp.Id)
	if err != nil {
		return nil, errors.New("ошибка обновления песни " + err.Error())
	}
	resultJson, err := json.Marshal(resp)
	if err != nil {
		return nil, errors.New("ошибка серилизации в json " + err.Error())
	}
	return resultJson, nil
}
