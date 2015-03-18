package util

import (
	"log"
)

type Student struct{
	name string
	first string
	second string
	}

func DormInit() ([]Student,[]string){
	dorms := []string{"Zeus", "Athena", "Hercules", "Bacchus", "Pluto"}
	
	student := make([]Student, len(dorms)*2)
	
	student[0].name = "Toby"
	student[0].first = "Bacchus"
	student[0].second = "Hercules"
	
	student[1].name = "Steve"
	student[1].first = "Zeus"
	student[1].second = "Pluto"
	
	student[2].name = "Andrea"
	student[2].first = "Athena"
	student[2].second = "Zeus"
	
	student[3].name = "Sarah"
	student[3].first = "Zeus"
	student[3].second = "Pluto"
	
	student[4].name = "Dave"
	student[4].first = "Athena"
	student[4].second = "Bacchus"
	
	student[5].name = "Jeff"
	student[5].first = "Hercules"
	student[5].second = "Pluto"
	
	student[6].name = "Fred"
	student[6].first = "Pluto"
	student[6].second = "Athena"
	
	student[7].name = "Suzie"
	student[7].first = "Bacchus"
	student[7].second = "Hercules"
	
	student[8].name = "Laura"
	student[8].first = "Bacchus"
	student[8].second = "Hercules"
	
	student[9].name = "Neil"
	student[9].first = "Hercules"
	student[9].second = "Athena"


	return student,dorms
}


func PrintSolution(s []Student, dorms []string, sol []int){
	
	var slots []int
	for index,_ := range dorms{
		slots = append(slots, index)
		slots = append(slots, index)
	} 
//	log.Println(slots)

	for index, _ := range sol{
		
		dorm := dorms[slots[index]]
		
		log.Printf("%s %s\n",s[index].name, dorm)
		
	}
	
	
	
}


func main() {
	student, dorms := DormInit()

	for _, value := range dorms {
		log.Println(value)
	}
	
	log.Println(student)
	
	sol := []int{0,0,0,0,0,0,0,0,0,0}
	
	PrintSolution(student,dorms,sol)

}
