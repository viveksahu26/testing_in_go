package eeee

import "log"

func PrintError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
