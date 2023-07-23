package main

import (
	"fmt"
	"scaffolder/helper"
	"strings"
)

type yamlVariableMap map[string]string

func (y *yamlVariableMap) String() string {
	output := []string{}

	for k, v := range *y {
		output = append(output, fmt.Sprintf("%s:%s", k, v))
	}
	return strings.Join(output, ",")

}

func (y *yamlVariableMap) Set(val string) error {
	values := strings.Split(val, ",")
	for _, v := range values {
		setVariable := strings.Split(v, ":")
		if len(setVariable) < 2 {
			helper.Fatal("invalid format for varibale assignment", false)
			return fmt.Errorf("invalid format for key-value pair: %s", v)
		}
		key := setVariable[0]
		valu := setVariable[1]
		(*y)[key] = valu

	}

	return nil
}
