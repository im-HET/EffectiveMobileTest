package utils

import (
	"errors"
	"fmt"
	"strconv"

	"strings"
	"time"
)

func ValidateString(v interface{}) (string, error) {
	if strings.HasPrefix(fmt.Sprint(v), "$$") && strings.HasSuffix(fmt.Sprint(v), "$$") {
		return fmt.Sprint(v), nil
	}
	return "$$" + fmt.Sprint(v) + "$$", nil
}

func ValidateDate(v interface{}) (string, error) {
	date, err := time.Parse("02.01.2006", fmt.Sprint(v))
	if err == nil {
		return "'" + date.String()[:10] + "'", nil
	}
	date, err = time.Parse("2006-01-02", fmt.Sprint(v))
	if err == nil {
		return "'" + date.String()[:10] + "'", nil
	}
	return "", errors.New("ожидается дата " + err.Error())
}
func ValidateInt(v interface{}) (string, error) {
	if _, err := strconv.Atoi(fmt.Sprint(v)); err != nil {
		return "", errors.New("ожидается целое число ")
	}
	return fmt.Sprint(v), nil
}
