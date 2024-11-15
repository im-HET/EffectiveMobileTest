package utils

import (
	"strings"
)

func GetColumns(params map[string]*Parameter) string {
	result := []string{}
	for _, p := range params {
		if len(p.values) > 0 {
			if p.Column == "" {
				result = append(result, p.Name)
			} else {
				result = append(result, p.Column)
			}
		}
	}
	return "(" + strings.Join(result, ", ") + ")"
}

func GetValues(params map[string]*Parameter) string {
	result := []string{}
	for _, p := range params {
		if len(p.values) > 0 && p.values[0] != "" {
			result = append(result, p.values[0])
		}
	}
	return "(" + strings.Join(result, ", ") + ")"
}
