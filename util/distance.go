package util

import (
	"log"
	"math"
	"sort"
)

type Items map[string]float64
type Dataset map[string]Items

type Sim_algorithm func(Dataset, string, string) float64

type Score struct {
	Point float64
	Name  string
	Movie string
}

func Initialize() Dataset {
	dataset := make(Dataset)

	m1 := make(Items)
	m1["Lady in the Water"] = 2.5
	m1["Snakes on a Plane"] = 3.5
	m1["Just My Luck"] = 3.0
	m1["Superman Returns"] = 3.5
	m1["You, Me and Dupree"] = 2.5
	m1["The Night Listener"] = 3.0
	dataset["Lisa Rose"] = m1

	m2 := make(Items)
	m2["Lady in the Water"] = 3.0
	m2["Snakes on a Plane"] = 3.5
	m2["Just My Luck"] = 1.5
	m2["Superman Returns"] = 5.0
	m2["You, Me and Dupree"] = 3.5
	m2["The Night Listener"] = 3.0
	dataset["Gene Seymour"] = m2

	m3 := make(Items)
	m3["Lady in the Water"] = 2.5
	m3["Snakes on a Plane"] = 3.0
	m3["Superman Returns"] = 3.5
	m3["The Night Listener"] = 4.0
	dataset["Michael Phillips"] = m3

	m4 := make(Items)
	m4["Snakes on a Plane"] = 3.5
	m4["Superman Returns"] = 4.0
	m4["Just My Luck"] = 3.0
	m4["The Night Listener"] = 4.5
	m4["You, Me and Dupree"] = 2.5
	dataset["Claudia Puig"] = m4

	m5 := make(Items)
	m5["Lady in the Water"] = 3.0
	m5["Snakes on a Plane"] = 4.0
	m5["Superman Returns"] = 3.0
	m5["Just My Luck"] = 2.0
	m5["The Night Listener"] = 3.0
	m5["You, Me and Dupree"] = 2.0
	dataset["Mick LaSalle"] = m5

	m6 := make(Items)
	m6["Lady in the Water"] = 3.0
	m6["Snakes on a Plane"] = 4.0
	m6["Superman Returns"] = 5.0
	m6["The Night Listener"] = 3.0
	m6["You, Me and Dupree"] = 3.5
	dataset["Jack Matthews"] = m6

	m7 := make(Items)
	m7["Snakes on a Plane"] = 4.5
	m7["Superman Returns"] = 4.0
	m7["You, Me and Dupree"] = 1.0
	dataset["Toby"] = m7

	return dataset

}

func Sim_distance(p Dataset, p1 string, p2 string) float64 {
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

func Sim_pearson(p Dataset, p1 string, p2 string) float64 {
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

			sum1Sq = sum1Sq + math.Pow(p[p1][key], 2)
			sum2Sq = sum2Sq + math.Pow(p[p2][key], 2)

			pSum = pSum + p[p1][key]*p[p2][key]

			n = n + 1
		}
	}
	if n == 0 {
		return 0.0
	}

	num := pSum - (sum1 * sum2 / n)
	den := math.Sqrt((sum1Sq - math.Pow(sum1, 2)/n) * (sum2Sq - math.Pow(sum2, 2)/n))
	if den == 0 {
		return 0.0
	}

	result = num / den

	return result
}

func reverseMap1(m map[float64]string) []Score {
	var result []Score
	var keys []float64
	for k := range m {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.Float64Slice(keys)))
	for _, k := range keys {
		var s Score
		s.Name = m[k]
		s.Point = k
		result = append(result, s)
	}
	//	log.Println(result)
	return result
}

func reverseMap2(m map[string]float64) []Score {
	m1 := make(map[float64]string)
	for k := range m {
		m1[m[k]] = k
	}
	var result []Score
	var keys []float64
	for k := range m1 {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.Float64Slice(keys)))
	for _, k := range keys {
		var s Score
		s.Movie = m1[k]
		s.Point = k
		result = append(result, s)
	}
	//	log.Println(result)
	return result
}

func TopMatches(p Dataset, name string, n int, f Sim_algorithm) []Score {
	var result []Score
	sim := make(map[float64]string)
	for key, _ := range p {
		if key != name {
			dis := f(p, name, key)
			sim[dis] = key
		}
	}
	result = reverseMap1(sim)

	if n < len(result) {
		result = result[:n]
	}

	return result
}

func GetRecommendations(p Dataset, name string, f Sim_algorithm) []Score {
	var result []Score
	totals := make(map[string]float64)
	simSums := make(map[string]float64)

	rankings := make(map[string]float64)

	for k := range p {
		if k != name {
			sim := f(p, name, k)
			if sim > 0 {
				for movie, _ := range p[k] {
					_, ok := p[name][movie]
					if ok != true {
						totals[movie] = totals[movie] + p[k][movie]*sim
						simSums[movie] = simSums[movie] + sim
					}
				}

			}
		}
	}

	for k := range totals {
		rankings[k] = totals[k] / simSums[k]
	}

	result = reverseMap2(rankings)

	return result
}

func TransformPrefs(p Dataset) Dataset {
	result := make(Dataset)

	for name, _ := range p {
		for movie, _ := range p[name] {
			_, ok := result[movie]
			if ok != true {
				result[movie] = make(Items)
			}
			result[movie][name] = p[name][movie]
		}
	}

	//	log.Println(result)

	return result
}

func CalculateSimilarItems(p Dataset, n int) map[string][]Score {
	result := make(map[string][]Score)

	d := TransformPrefs(p)

	for k := range d {
		scores := TopMatches(d, k, n, Sim_distance)
		result[k] = scores
	}

	log.Println(result["Lady in the Water"])
	log.Println(result["Snakes on a Plane"])

	return result
}

func GetRecommendationItems(p Dataset, sim map[string][]Score, name string){
	userRatings := p[name]

	scores := make(map[string]float64)
	totalSums := make(map[string]float64)

	for movie, rating := range userRatings {
		for _, item := range sim[movie] {
			_, ok := userRatings[item.Name]
			if ok != true {
				scores[item.Name] = scores[item.Name] + item.Point*rating
				totalSums[item.Name] = totalSums[item.Name] + item.Point
			}
		}
	}
	
	rankings := make(map[float64]string)
	for k := range scores {
		rankings[scores[k]/totalSums[k]] = k
	}
	result := reverseMap1(rankings)
	
	log.Println(result)
}
