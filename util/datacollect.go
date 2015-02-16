package util

import (
	"bufio"
	"os"
	"log"
	"strings"
	"strconv"
)



type MovieTitle map[string]string

func GetFullPath(filename string)string{
	pwd,err := os.Getwd()
	if err != nil{
		log.Fatal(err)
	}
	result := pwd + "\\data\\movielens\\" + filename
	return result
}

func GetScheduleFilepath()string{
	pwd,err := os.Getwd()
	if err != nil{
		log.Fatal(err)
	}
	result := pwd + "\\data\\flight\\schedule.txt"
	return result
}

func LoadMovieTitles()MovieTitle{
	result := make(MovieTitle)
	
	itemfile := GetFullPath("u.item")
	
	f,err := os.Open(itemfile)
	if err != nil{
		log.Fatal(err)
	}
	defer f.Close()

	r := bufio.NewScanner(f)
	r.Split(bufio.ScanLines)
	for r.Scan(){
		s := strings.Split(r.Text(), "|")[:2]
		result[s[0]] = s[1] 
	}
//	log.Println(len(result))
	return result
}

func LoadMovieData(movie MovieTitle)Dataset{
	result := make(Dataset)
	
	datafile := GetFullPath("u.data")
	f,err := os.Open(datafile)
	if err != nil{
		log.Fatal(err)
	}
	defer f.Close()
	
	r := bufio.NewScanner(f)
	r.Split(bufio.ScanLines)
	for r.Scan(){
		s := strings.Split(r.Text(), "\t")
//		log.Println(s)
		userid := s[0]
		movieid := s[1]
		rating := s[2]
		val, err := strconv.ParseFloat(rating,64)
		if err != nil{
			log.Println(err.Error())
		}
		item,ok := result[userid]
		if ok != true{
			item = make(Items)
		}
		item[movie[movieid]] = val
		result[userid] = item 
	}
	return result
}

func LoadScheduleData()Flight{
	datafile := GetScheduleFilepath()
	f,err := os.Open(datafile)
	if err != nil{
		log.Fatal(err)
	}
	defer f.Close()
	
	result := make(Flight)
	r := bufio.NewScanner(f)
	r.Split(bufio.ScanLines)
	for r.Scan(){
		line := strings.Split(r.Text(), ",")
//		log.Println(s)

		addr := line[0] + "-" + line[1]
		
		
		var s Schedule
//		s.Origin = line[0]
//		s.Dest = line[1]
		s.Depart = line[2]
		s.Arrive = line[3]
		s.Price,_ = strconv.Atoi(line[4])
		
		result[addr] = append(result[addr],s)
	}
//	log.Println(result)
	return result 
} 





