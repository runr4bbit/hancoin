package utils

import "log"

func HandleFunc(err error) {
	if err != nil {
		log.Panic(err)
	}
}
