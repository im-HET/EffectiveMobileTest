package utils

import (
	"errors"
	"strings"
)

type Parameter struct {
	Name      string
	values    []string
	Condition string
	Column    string
	Required  bool
	Validator func(value interface{}) (string, error)
}

func (p Parameter) String() string {
	resultStr := []string{}
	if len(p.values) == 0 {
		return ""
	}
	if p.Condition == "" {
		return p.Name + " " + p.values[0]
	}
	for _, v := range p.values {
		parameterStr := []string{}
		if p.Column != "" {
			parameterStr = append(parameterStr, p.Column)
		} else {
			parameterStr = append(parameterStr, p.Name)
		}
		parameterStr = append(parameterStr, p.Condition)
		parameterStr = append(parameterStr, v)
		resultStr = append(resultStr, strings.Join(parameterStr, " "))
	}
	return "( " + strings.Join(resultStr, " or ") + " )"
}

func (p *Parameter) AddValue(value interface{}) error {
	s, err := p.Validator(value)
	if err != nil {
		return errors.New("не верное значение " + err.Error())
	}
	p.values = append(p.values, s)
	return nil
}

func (p Parameter) GetValue() string {
	if len(p.values) == 0 {
		return ""
	}
	return p.values[0]
}

func (p *Parameter) Clear() {
	p.values = []string{}
}
