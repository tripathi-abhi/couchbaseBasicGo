package utils

import "log"

func CheckErrorNil(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
