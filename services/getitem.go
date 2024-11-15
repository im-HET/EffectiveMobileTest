package services

import (
	"encoding/json"
	"errors"
	"mediaLibrary_v2/db"

	"mediaLibrary_v2/utils"
	"strings"
)

type GetItem struct {
	service
}

func NewGetItem() GetItem {
	i := GetItem{}
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

func (i GetItem) Get() ([]byte, error) {
	type responseBeforeJson struct { //используем указатели для того
		Id          *int    `db:"id"`   //что бы без ошибок обрабатывать null значения из базы
		Song        *string `db:"song"` //указываем тэг для работы sqlx
		Group       *string `db:"name"`
		ReleaseDate *string `db:"releasedate"`
		Link        *string `db:"link"`
		Text        *string `db:"text"`
	}
	sql := []string{
		"SELECT",
		"library.id,",
		"library.song,",
		"groups.name,",
		"to_char(library.releasedate, 'dd.mm.yyyy') as releasedate,",
		"library.link,",
		"library.text",
		"FROM library, groups",
		utils.GetWhere(
			"( library.groupid = groups.id )",
			*i.expectedParameters["id"],
		)}
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
