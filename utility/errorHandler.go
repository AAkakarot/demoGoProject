package utility

import "log"

func ErrorHandler(err error , msg string) {
	if err != nil{
		log.Fatalf("%s: %s", msg, err)
	}
}

