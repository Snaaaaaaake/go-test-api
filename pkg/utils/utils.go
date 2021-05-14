package utils

import "github.com/sirupsen/logrus"

func ErrorCheck(err error, text ...string) {
	if err != nil {
		if len(text) > 0 {
			logrus.Fatalf(text[0]+": ", err.Error())
		} else {
			logrus.Fatal(err)
		}
	}
}
