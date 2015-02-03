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
	"sort"
)

type Movies map[string]float64
type Persons map[string]Movies

type sim_algorithm func(Persons,string,string) float64

type Score struct{
	Point float64
	Name string
}

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
			result = result + math.Pow(p[p1][key]-p[p2][key], 2)
		}
	}

	result = 1 / (1 + result)

	return result
}

func sim_pearson(p Persons, p1 string, p2 string) float64 {
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
	sum1 := 0.0
	sum2 := 0.0
	sum1Sq := 0.0
	sum2Sq := 0.0
	pSum := 0.0
	n := 0.0
	for key, _ := range p[p1] {
		_, ok := p[p2][key]
		if ok {
			sum1 = sum1 + p[p1][key]
			sum2 = sum2 + p[p2][key]
			
			sum1Sq = sum1Sq + math.Pow(p[p1][key],2)
			sum2Sq = sum2Sq + math.Pow(p[p2][key],2)
			
			pSum = pSum + p[p1][key]*p[p2][key]
			
			n = n+1
		}
	}
	if n == 0 {
		return 0.0
	}
	
	num := pSum - (sum1*sum2/n)
	den := math.Sqrt((sum1Sq-math.Pow(sum1,2)/n)*(sum2Sq-math.Pow(sum2,2)/n))
	if den == 0 {
		return 0.0
	}
	
	result = num/den

	return result
}

func reverseMap(m map[float64]string)[]Score{
	var result []Score
	var keys []float64
	for k := range m{
		keys = append(keys,k)
	}
	log.Println(keys)
	sort.Sort(sort.Reverse(sort.Float64Slice(keys)))
	log.Println(keys)
	for _,k := range keys{
		var s Score
		s.Name = m[k]
		s.Point = k
		result = append(result,s)
	}
	log.Println(result)
	return result
}

func topMatches(p Persons, name string, n int, f sim_algorithm)[]Score{
	result := make([]Score,len(p))
	sim := make(map[float64]string)
	for key,_:=range p{
		if key != name{
			dis := f(p, name, key)
			sim[dis] = key
		}
	}
	result = reverseMap(sim)
	
	return result[:n]
}

func main() {

	d := Initialize()

//	log.Println(sim_distance(d, "test", "Jack Matthews"))
	log.Println(sim_distance(d, "Lisa Rose", "Gene Seymour"))
	log.Println(sim_pearson(d, "Lisa Rose", "Gene Seymour"))
	log.Println(sim_pearson(d, "Toby", "Claudia Puig"))
	
	log.Println(topMatches(d, "Toby", 3, sim_pearson))

}
