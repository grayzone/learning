package util

import (
	"log"
	
)

type Schedule struct {
	Name string
	Origin string
	Dest string
	Depart string
	Arrive string
	Price int
}

type Flight map[string][]Schedule

const DESTINATION = "LGA"

func InitOpt()[]Schedule{
	var result []Schedule
	
	var s1 Schedule
	s1.Name = "Seymour"
	s1.Origin = "BOS"
	result = append(result, s1)
	
	s2 := Schedule{Name:"Franny",Origin:"DAL"}
	result = append(result, s2)
	
	s3 := Schedule{Name:"Zooey",Origin:"CAK"}
	result = append(result, s3)
	
	s4 := Schedule{Name:"Walt",Origin:"MIA"}
	result = append(result, s4)
	
	s5 := Schedule{Name:"Buddy",Origin:"ORD"}
	result = append(result, s5)
	
	s6 := Schedule{Name:"Les",Origin:"OMA"}
	result = append(result, s6)
	
//	log.Println(result)

	return result
}


func PrintSchedule(p []Schedule, s []int){
	
	f := LoadScheduleData()
	for index, value := range p {
		name := value.Name
		origin := value.Origin
//		log.Println(f[origin + "-" + DESTINATION])
//		log.Println(f[DESTINATION + "-" + origin])
		out := f[origin + "-" + DESTINATION][s[2*index]]
		ret := f[DESTINATION + "-" + origin][s[2*index+1]]
		log.Printf("%s:%s %v %v\n", name,origin, out, ret)
	}
}


func ScheduleCost(p []Schedule, s []int){
	
	totalprice := 0
	latestarrival := 0
	earliestdep := 1440
	
	
	f := LoadScheduleData()
	for index, value := range p{
		origin := value.Origin
		out := f[origin + "-" + DESTINATION][s[2*index]]
		ret := f[DESTINATION + "-" + origin][s[2*index+1]]
		
		totalprice = totalprice + out.Price
		totalprice = totalprice + ret.Price
		
		if latestarrival < GetMinutes(out.Arrive){
			latestarrival = GetMinutes(out.Arrive)
		}
		if earliestdep > GetMinutes(ret.Depart){
			earliestdep = GetMinutes(ret.Depart)
		}
	}
	log.Println(totalprice)
	log.Println(latestarrival)
	log.Println(earliestdep)
	
	totalwait := 0
	for index, value := range p{
		origin := value.Origin
		out := f[origin + "-" + DESTINATION][s[2*index]]
		ret := f[DESTINATION + "-" + origin][s[2*index+1]]
		
		totalwait = latestarrival -  + out.Price
		totalprice = totalprice + ret.Price
		
		if latestarrival < GetMinutes(out.Arrive){
			latestarrival = GetMinutes(out.Arrive)
		}
		if earliestdep > GetMinutes(ret.Depart){
			earliestdep = GetMinutes(ret.Depart)
		}
	}
	
	
	
	
	
}