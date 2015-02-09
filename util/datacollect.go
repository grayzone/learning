package util

import (
	"bufio"
	"os"
	"log"
	"strings"
	"strconv"

)

type MovieTitle map[int]string

func GetFullPath(filename string)string{
	pwd,err := os.Getwd()
	if err != nil{
		log.Fatal(err)
	}
	result := pwd + "\\data\\movielens\\" + filename
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
		index, err := strconv.Atoi(s[0])
		if err != nil{
			log.Println(err.Error())
		}else{
			result[index] = s[1] 
		}
	}
	log.Println(len(result))
	return result
}

func LoadMovieLens(){
	
	title := LoadMovieTitles()
	log.Println(title)
	
}

