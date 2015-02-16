package util

import (
	"fmt"
	"log"

)

func GetMinutes(s string)int{
	var hour int
	var minute int
	_, err := fmt.Sscanf(s, "%d:%d",&hour,&minute)
	if err != nil{
		log.Println(err.Error())
		return 0
	}
	return hour*60 + minute
}