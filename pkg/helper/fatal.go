package helper

import (
	"fmt"
	"os"
)

/*
Helper function to handle errors

message is the error message to be sent,

checkErr is if we should check err variable passed in err argument, like

	if err != nil {
		...
	   }
*/
func Fatal(message string, checkErr bool, err ...error) {
	if !checkErr {
		fmt.Println(message)
		os.Exit(1)
	} else {
		if len(err) > 0 && err[0] != nil {
			fmt.Println(message)
			os.Exit(1)
		}
	}
}
