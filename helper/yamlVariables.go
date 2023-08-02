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
		setVariable := strings.SplitN(v, ":", 2)
		if len(setVariable) < 2 {
			Fatal("Invalid format for variable assignment", false)
			return fmt.Errorf("invalid format for key-value pair: %s", v)
		}
		key := strings.TrimSpace(setVariable[0])
		value := strings.TrimSpace(setVariable[1])
		(*y)[key] = value

	}

	return nil
}
