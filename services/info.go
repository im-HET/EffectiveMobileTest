package services

import (
	"encoding/json"
	"errors"
	"mediaLibrary_v2/db"
	"mediaLibrary_v2/utils"
	"strings"
)

type Info struct {
	service
}

func NewInfo() Info {
	i := Info{}
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

	i.expectedParameters = map[string]*utils.Parameter{
		song.Name:  &song,
		group.Name: &group,
	}
	return i
}

func (i Info) Get() ([]byte, error) {
	type responseBeforeJson struct {
		ReleaseDate *string `db:"releasedate"` //Указатели что бы без ошибок обрабатывать null значения из базы
		Link        *string `db:"link"`        //указываем тэг для работы sqlx
		Text        *string `db:"text"`
	}
	sql := []string{
		"SELECT",
		"to_char(library.releasedate, 'dd.mm.yyyy') as releasedate,",
		"library.link,",
		"library.text",
		"FROM library, groups",
		utils.GetWhere(
			"( library.groupid = groups.id )",
			*i.expectedParameters["song"],
			*i.expectedParameters["group"],
		)}
	resp := responseBeforeJson{}
	err := db.SqlGet(strings.Join(sql, " "), &resp)
	if err != nil {
		if errors.Is(err, db.ErrNoRows{}) {
			return nil, errors.New("нет в базе " + err.Error())
		} else {
			return nil, errors.New(err.Error())
		}
	}
	resultJson, err := json.Marshal(resp)
	if err != nil {
		return nil, errors.New("ошибка серилизации в json " + err.Error())
	}
	return resultJson, nil
}
