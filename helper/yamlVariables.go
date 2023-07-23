package helper

import (
	"fmt"
	"strings"
)

type YamlVariableMap map[string]string

func (y *YamlVariableMap) String() string {
	output := []string{}

	for k, v := range *y {
		output = append(output, fmt.Sprintf("%s:%s", k, v))
	}
	return strings.Join(output, ",")

}

func (y *YamlVariableMap) Set(val string) error {
	values := strings.Split(val, ",")
	for _, v := range values {
		setVariable := strings.Split(v, ":")
		if len(setVariable) < 2 {
			Fatal("Invalid format for varibale assignment", false)
			return fmt.Errorf("Invalid format for key-value pair: %s", v)
		}
		key := setVariable[0]
		value := setVariable[1]
		(*y)[key] = value

	}

	return nil
}
