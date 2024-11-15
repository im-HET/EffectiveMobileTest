package services

import (
	"encoding/json"
	"errors"
	"mediaLibrary_v2/db"
	"mediaLibrary_v2/utils"
	"strconv"
	"strings"
)

type GetItems struct {
	service
}

func NewGetItems() GetItems { //Конструктор для сервиса выдачи списка песен
	i := GetItems{}
	song := utils.Parameter{ //Определяем ожидаемые/возможные параметры для данного сервиса
		Name:      "song",
		Required:  false,                //Если наличие параметра обязательно для работы сервиса ставим true
		Condition: " LIKE ",             //операция сравнения для использования параметра в конструкции where
		Validator: utils.ValidateString, //Валидатор параметра. можно добавить свой или выбрать "стандартный"
		Column:    "library.song",
	}
	group := utils.Parameter{
		Name:      "group",
		Required:  false,
		Condition: " LIKE ",
		Column:    "groups.name",
		Validator: utils.ValidateString,
	}
	id := utils.Parameter{
		Name:      "id",
		Required:  false,
		Condition: " = ",
		Column:    "library.id",
		Validator: utils.ValidateInt,
	}
	dataStart := utils.Parameter{
		Name:      "datastart",
		Required:  false,
		Condition: " >= ",
		Column:    "library.releasedate",
		Validator: utils.ValidateDate,
	}
	dataEnd := utils.Parameter{
		Name:      "dataend",
		Required:  false,
		Condition: " <= ",
		Column:    "library.releasedate",
		Validator: utils.ValidateDate,
	}
	link := utils.Parameter{
		Name:      "link",
		Required:  false,
		Condition: " LIKE ",
		Column:    "library.link",
		Validator: utils.ValidateString,
	}
	limit := utils.Parameter{
		Name:      "limit",
		Required:  false,
		Validator: utils.ValidateInt,
	}
	offset := utils.Parameter{
		Name:      "offset",
		Required:  false,
		Validator: utils.ValidateInt,
	}

	i.expectedParameters = map[string]*utils.Parameter{
		id.Name:        &id,
		song.Name:      &song,
		group.Name:     &group,
		dataStart.Name: &dataStart,
		dataEnd.Name:   &dataEnd,
		link.Name:      &link,
		limit.Name:     &limit,
		offset.Name:    &offset,
	}

	return i
}

func (i GetItems) Get() ([]byte, error) {
	type responseBeforeJson struct {
		Name        string
		NumOfPages  int
		CurrentPage int
		Items       []struct {
			Id          *int    `db:"id"`   //что бы без ошибок обрабатывать null значения из базы
			Song        *string `db:"song"` //указываем тэг для работы sqlx
			Group       *string `db:"name"`
			ReleaseDate *string `db:"releasedate"`
			Link        *string `db:"link"`
		}
	}

	//Формируем запрос количества записей для вычисления пагинации
	countSql := []string{
		"SELECT count(library.id) FROM library, groups",
		utils.GetWhere(
			"( library.groupid = groups.id )",
			*i.expectedParameters["id"],
			*i.expectedParameters["song"],
			*i.expectedParameters["group"],
			*i.expectedParameters["datastart"],
			*i.expectedParameters["dataend"],
			*i.expectedParameters["link"],
		),
	}
	//Формируем sql запрос
	selectSql := []string{
		"SELECT",
		"library.id,",
		"library.song,",
		"groups.name,",
		"to_char(library.releasedate, 'dd.mm.yyyy') as releasedate,",
		"library.link",
		"FROM library, groups",
		utils.GetWhere(
			"( library.groupid = groups.id )",
			*i.expectedParameters["id"],
			*i.expectedParameters["song"],
			*i.expectedParameters["group"],
			*i.expectedParameters["datastart"],
			*i.expectedParameters["dataend"],
			*i.expectedParameters["link"],
		),
		(*i.expectedParameters["limit"]).String(),
		(*i.expectedParameters["offset"]).String(),
	}

	countRows := 0
	//Получаем количество записей для того что бы узнать количество страниц пагинации
	err := db.SqlGet(strings.Join(countSql, " "), &countRows)
	if err != nil {
		return nil, errors.New("ошибка подсчета количества записей " + err.Error())
	}

	//Задаем имя ответа
	resp := responseBeforeJson{Name: "mediaElements"}
	//Получаем записи
	err = db.SqlSelect(strings.Join(selectSql, " "), &resp.Items)
	if err != nil {
		return nil, errors.New("ошибка получения записей " + err.Error())
	}
	//Добавляем в ответ количество страниц
	pLimit, _ := i.expectedParameters["limit"]
	limit, _ := strconv.Atoi(pLimit.GetValue())
	//Добавляем в ответ текущую страницу
	pOffset, _ := i.expectedParameters["offset"]
	offset, _ := strconv.Atoi(pOffset.GetValue())
	if limit > 0 {
		resp.NumOfPages = countRows/limit + 1
		resp.CurrentPage = offset/limit + 1
	} else {
		resp.NumOfPages = 1
		resp.CurrentPage = 1
	}
	resultJson, err := json.Marshal(resp)
	if err != nil {
		return nil, errors.New("ошибка серилизации в json " + err.Error())
	}
	return resultJson, nil
}
