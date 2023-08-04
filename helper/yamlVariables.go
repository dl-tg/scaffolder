package helper

import (
	"fmt"
	"strings"
)

// YamlVariableMap is a type that represents a map of string variables in YAML format.
type YamlVariableMap map[string]string

// String returns a string representation of the YamlVariableMap.
// The string will be in the format "key1:value1,key2:value2,..."
func (y *YamlVariableMap) String() string {
	output := []string{}

	for k, v := range *y {
		output = append(output, fmt.Sprintf("%s:%s", k, v))
	}
	return strings.Join(output, ",")
}

// Set parses the input string and sets the variables in the YamlVariableMap.
// The input string should be in the format "key1:value1,key2:value2,...".
func (y *YamlVariableMap) Set(val string) error {
	values := strings.Split(val, ",")
	for _, v := range values {
		setVariable := strings.SplitN(v, ":", 2)
		if len(setVariable) < 2 {
			Fatal("Invalid format for variable assignment", false) // Assuming "Fatal" is a function to handle fatal errors
			return fmt.Errorf("invalid format for key-value pair: %s", v)
		}
		key := strings.TrimSpace(setVariable[0])
		value := strings.TrimSpace(setVariable[1])
		(*y)[key] = value
	}

	return nil
}
