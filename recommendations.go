// A dictionary of movie critics and their ratings of a small
// set of movies

/*
critics :={'Lisa Rose': {'Lady in the Water': 2.5, 'Snakes on a Plane': 3.5,
'Just My Luck': 3.0, 'Superman Returns': 3.5, 'You, Me and Dupree': 2.5,
'The Night Listener': 3.0},
'Gene Seymour': {'Lady in the Water': 3.0, 'Snakes on a Plane': 3.5,
'Just My Luck': 1.5, 'Superman Returns': 5.0, 'The Night Listener': 3.0,
'You, Me and Dupree': 3.5},
'Michael Phillips': {'Lady in the Water': 2.5, 'Snakes on a Plane': 3.0,
'Superman Returns': 3.5, 'The Night Listener': 4.0},
'Claudia Puig': {'Snakes on a Plane': 3.5, 'Just My Luck': 3.0,
'The Night Listener': 4.5, 'Superman Returns': 4.0,
'You, Me and Dupree': 2.5},
'Mick LaSalle': {'Lady in the Water': 3.0, 'Snakes on a Plane': 4.0,
'Just My Luck': 2.0, 'Superman Returns': 3.0, 'The Night Listener': 3.0,
'You, Me and Dupree': 2.0},
'Jack Matthews': {'Lady in the Water': 3.0, 'Snakes on a Plane': 4.0,
'The Night Listener': 3.0, 'Superman Returns': 5.0, 'You, Me and Dupree': 3.5},
'Toby': {'Snakes on a Plane':4.5,'You, Me and Dupree':1.0,'Superman Returns':4.0}}

*/
package main

import (
	"log"
	"math"
)

type Movies map[string]float64
type Persons map[string]Movies

func Initialize() Persons {
	dataset := make(Persons)

	m1 := make(Movies)
	m1["Lady in the Water"] = 2.5
	m1["Snakes on a Plane"] = 3.5
	m1["Just My Luck"] = 3.0
	m1["Superman Returns"] = 3.5
	m1["You, Me and Dupree"] = 2.5
	m1["The Night Listener"] = 3.0
	dataset["Lisa Rose"] = m1

	m2 := make(Movies)
	m2["Lady in the Water"] = 3.0
	m2["Snakes on a Plane"] = 3.5
	m2["Just My Luck"] = 1.5
	m2["Superman Returns"] = 5.0
	m2["You, Me and Dupree"] = 3.5
	m2["The Night Listener"] = 3.0
	dataset["Gene Seymour"] = m2

	m3 := make(Movies)
	m3["Lady in the Water"] = 2.5
	m3["Snakes on a Plane"] = 3.0
	m3["Superman Returns"] = 3.5
	m3["The Night Listener"] = 4.0
	dataset["Michael Phillips"] = m3

	m4 := make(Movies)
	m4["Lady in the Water"] = 2.5
	m4["Snakes on a Plane"] = 3.5
	m4["Superman Returns"] = 4.0
	m4["Just My Luck"] = 3.0
	m4["The Night Listener"] = 4.5
	m4["You, Me and Dupree"] = 2.0
	dataset["Claudia Puig"] = m3

	m5 := make(Movies)
	m5["Lady in the Water"] = 3.0
	m5["Snakes on a Plane"] = 4.0
	m5["Superman Returns"] = 3.0
	m5["Just My Luck"] = 2.0
	m5["The Night Listener"] = 3.0
	m5["You, Me and Dupree"] = 2.0
	dataset["Mick LaSalle"] = m5

	m6 := make(Movies)
	m6["Lady in the Water"] = 3.0
	m6["Snakes on a Plane"] = 4.0
	m6["Superman Returns"] = 5.0
	m6["The Night Listener"] = 3.0
	m6["You, Me and Dupree"] = 3.5
	dataset["Jack Matthews"] = m6

	m7 := make(Movies)
	m7["Snakes on a Plane"] = 4.5
	m7["Superman Returns"] = 4.0
	m7["You, Me and Dupree"] = 1.0
	dataset["Toby"] = m7

	return dataset

}

func sim_distance(p Persons, p1 string, p2 string) float64 {
	result := 0.0

	_, ok := p[p1]
	if ok != true {
		log.Println(p1 + " is not existing in the dataset.")
		return 0.0
	}
	_, ok = p[p2]
	if ok != true {
		log.Println(p2 + " is not existing in the dataset.")
		return 0.0
	}
	for key, _ := range p[p1] {
		_, ok := p[p2][key]
		if ok {
			result = result + math.Pow(p[p1][key]-p[p2][key],2)
		}
	}
	
	result = 1 /(1+ result)
	
	log.Println(result)

	return result
}

func main() {

	d := Initialize()

	sim_distance(d, "test", "Jack Matthews")
	sim_distance(d, "Lisa Rose", "Gene Seymour")

	a := math.Sqrt(math.Pow(5-4, 2) + math.Pow(4-1, 2))
	log.Println(a)

	b := 1 / (1 + math.Sqrt(math.Pow(5-4, 2)+math.Pow(4-1, 2)))
	log.Println(b)

}
