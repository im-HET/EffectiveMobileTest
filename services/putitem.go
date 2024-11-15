package services

import (
	"encoding/json"
	"errors"
	"mediaLibrary_v2/db"
	"mediaLibrary_v2/utils"
	"strconv"
	"strings"
)

type PutItem struct {
	service
}

func NewPutItem() PutItem {
	p := PutItem{}
	song := utils.Parameter{ //описываем ожидаемые параметры
		Name:      "song",
		Required:  true,
		Condition: " = ",
		Column:    "library.song",
		Validator: utils.ValidateString,
	}
	group := utils.Parameter{
		Name:      "group",
		Required:  true,
		Condition: " = ",
		Column:    "groups.name",
		Validator: utils.ValidateString,
	}
	p.expectedParameters = map[string]*utils.Parameter{
		song.Name:  &song,
		group.Name: &group,
	}
	return p
}

func (p PutItem) Put() ([]byte, error) {
	type responseBeforeJson struct {
		Id *int `db:"id"`
	}
	pgroup := NewPutGroup()
	pg := pgroup.expectedParameters["group"]
	pg.AddValue(p.expectedParameters["group"].GetValue())
	idGroup := responseBeforeJson{}
	b, err := pgroup.Put()
	if err != nil {
		return nil, err
	}
	json.Unmarshal(b, &idGroup)

	sql := []string{
		"SELECT",
		"library.id",
		"FROM library, groups",
		utils.GetWhere(
			"groups.id = "+strconv.Itoa(*idGroup.Id),
			*p.expectedParameters["song"],
		),
	}
	var songId int = 0
	err = db.SqlGet(strings.Join(sql, " "), &songId)
	if err != nil {
		if errors.Is(err, db.ErrNoRows{}) { //если песня не была найдена по имени то добавляем новую
			sql = []string{
				"INSERT INTO",
				"library",
				"(song, groupid)",
				"VALUES (",
				(*p.expectedParameters["song"]).GetValue(),
				", ",
				strconv.Itoa(*idGroup.Id),
				") RETURNING id",
			}
			err = db.SqlExec(strings.Join(sql, " "), &songId)
			if err != nil {
				return nil, errors.New("ошибка добавления песни " + err.Error())
			}
		} else {
			return nil, errors.New("ошибка получения id песни " + err.Error())
		}
	}
	resp := responseBeforeJson{Id: &songId}
	resultJson, err := json.Marshal(resp)
	if err != nil {
		return nil, errors.New("ошибка серилизации в json " + err.Error())
	}
	return resultJson, nil
}
