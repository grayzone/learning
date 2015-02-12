package util

import (
	"log"
)

func InitOpt() {
	people := make(map[string]string)
	people["Seymour"] = "BOS"
	people["Franny"] = "DAL"
	people["Zooey"] = "CAK"
	people["Walt"] = "MIA"
	people["Buddy"] = "ORD"
	people["Les"] = "OMA"

	destination := "LGA"

	log.Println(people)
	log.Println(destination)
	

}


