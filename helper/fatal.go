package helper

import (
	"fmt"
	"os"
)

/*
Fatal is a helper function to handle errors and exit the program if necessary.

Parameters:
  - message: The error message to be displayed when an error occurs.
  - checkErr: A boolean flag to determine if the function should check the err argument.
    If checkErr is false, the function will print the error message and exit.
  - err: An optional variadic argument of type error. If provided, the function will check if it's not nil.

Example Usage:

	err := SomeFunction()
	Fatal("Error occurred while executing SomeFunction.", true, err)
	// Result: If 'err' is not nil, the error message will be displayed, and the program will exit with status code 1.

	Fatal("Error occurred during critical operation.", false)
	// Result: The error message will be displayed, and the program will exit with status code 1.

	Fatal("This will not exit the program.", false, nil)
	// Result: The error message will be displayed, but the program will not exit since 'err' is nil.
*/
func Fatal(message string, checkErr bool, err ...error) {
	if !checkErr || (checkErr && ErrorsExists(err...)) {
		fmt.Println(message)
		os.Exit(1)
	}
}

func ErrorsExists(errors ...error) bool {
	return len(errors) > 0 && errors[0] != nil
}
