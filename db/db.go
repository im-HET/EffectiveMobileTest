package db

import (
	"errors"
	"fmt"
	"os"

	"mediaLibrary_v2/settings"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
)

type ErrNoRows struct {
}

func (e ErrNoRows) Error() string {
	return "no rows"
}

var db *sqlx.DB

func StartDB() {
	connectionString := []string{
		"host=localhost",
		"sslmode=disable",
		"user=postgres",
		"password=Ytngfhjkz_1",
		//"dbname=mediaLibraryDB",
		//"user=postgres",
		//"password=postgres",
	}
	var err error
	db, err = sqlx.Connect("postgres", strings.Join(connectionString, " "))
	if err != nil {
		fmt.Println("ошибка открытия бд: ", err)
		os.Exit(1)
	}
	_, err = db.Exec("CREATE DATABASE \"" + settings.DBname + "\" WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'Russian_Russia.1251'")
	if err != nil {
		fmt.Println(err)
	}
	db.Close()
	fmt.Println(strings.Join(append(connectionString, "dbname="+settings.DBname), " "))
	db, err = sqlx.Connect("postgres", strings.Join(append(connectionString, "dbname="+settings.DBname), " "))
	if err != nil {
		fmt.Println("не удалось выбрать базу ", err)
	}
	instanse, err := postgres.WithInstance(db.DB, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance("file://migrations", settings.DBname, instanse)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = m.Up()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	db.Close()

	db, err = sqlx.Connect("postgres", strings.Join(append(connectionString, "dbname=mediaLibraryDB"), " "))
	if err != nil {
		fmt.Println("не удалось выбрать базу ", err.Error())
	}

	fmt.Println("Соединение с БД установлено")
}

func SqlSelect(sqlStr string, v any) error {
	fmt.Println(sqlStr)
	err := db.Select(v, sqlStr)
	if err != nil {
		return errors.New("ошибка получения набора из бд " + err.Error())
	}
	return nil
}

func SqlGet(sqlStr string, v any) error {
	fmt.Println(sqlStr)
	err := db.Get(v, sqlStr)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return ErrNoRows{}
		}
		return errors.New("ошибка получения значения из бд " + err.Error())
	}
	return nil
}
func SqlExec(sqlStr string, v any) error {
	fmt.Println(sqlStr)
	err := db.QueryRow(sqlStr).Scan(v)
	if err != nil {
		return errors.New("ошибка записи значения в бд " + err.Error())
	}
	return nil
}
