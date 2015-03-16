package util

import (
	"log"
	"math"
	"math/rand"
	"time"
)

type Schedule struct {
	Name   string
	Origin string
	Dest   string
	Depart string
	Arrive string
	Price  int
}

type Flight map[string][]Schedule

const DESTINATION = "LGA"

func InitOpt() []Schedule {
	var result []Schedule

	var s1 Schedule
	s1.Name = "Seymour"
	s1.Origin = "BOS"
	result = append(result, s1)

	s2 := Schedule{Name: "Franny", Origin: "DAL"}
	result = append(result, s2)

	s3 := Schedule{Name: "Zooey", Origin: "CAK"}
	result = append(result, s3)

	s4 := Schedule{Name: "Walt", Origin: "MIA"}
	result = append(result, s4)

	s5 := Schedule{Name: "Buddy", Origin: "ORD"}
	result = append(result, s5)

	s6 := Schedule{Name: "Les", Origin: "OMA"}
	result = append(result, s6)

	//	log.Println(result)

	return result
}

func PrintSchedule(p []Schedule, s []int) {

	f := LoadScheduleData()
	for index, value := range p {
		name := value.Name
		origin := value.Origin
		//		log.Println(f[origin + "-" + DESTINATION])
		//		log.Println(f[DESTINATION + "-" + origin])
		out := f[origin+"-"+DESTINATION][s[2*index]]
		ret := f[DESTINATION+"-"+origin][s[2*index+1]]
		log.Printf("%s:%s %v %v\n", name, origin, out, ret)
	}
}

func ScheduleCost(p []Schedule, s []int) int {

	totalprice := 0
	latestarrival := 0
	earliestdep := 1440

	f := LoadScheduleData()
	for index, value := range p {
		origin := value.Origin
		out := f[origin+"-"+DESTINATION][s[2*index]]
		ret := f[DESTINATION+"-"+origin][s[2*index+1]]

		totalprice = totalprice + out.Price
		totalprice = totalprice + ret.Price

		if latestarrival < GetMinutes(out.Arrive) {
			latestarrival = GetMinutes(out.Arrive)
		}
		if earliestdep > GetMinutes(ret.Depart) {
			earliestdep = GetMinutes(ret.Depart)
		}
	}
	//	log.Println(totalprice)
	//	log.Println(latestarrival)
	//	log.Println(earliestdep)

	totalwait := 0
	for index, value := range p {
		origin := value.Origin
		out := f[origin+"-"+DESTINATION][s[2*index]]
		ret := f[DESTINATION+"-"+origin][s[2*index+1]]

		totalwait = totalwait + latestarrival - GetMinutes(out.Depart)
		totalwait = totalwait + GetMinutes(ret.Arrive) - earliestdep

		if latestarrival > earliestdep {
			totalprice = totalprice + 50
		}

	}

	return totalprice + totalwait
}

func RandomOptimize(p []Schedule) {
	best := 99999999
	var bestr []int
	rand.Seed(time.Now().UnixNano())

	l := len(p)
	for i := 0; i < 1000; i++ {
		var seed []int
		for j := 0; j < l; j++ {
			seed = append(seed, rand.Intn(10))
			seed = append(seed, rand.Intn(10))
		}
		//		log.Println(seed)
		cost := ScheduleCost(p, seed)
		if cost < best {
			best = cost
			bestr = seed
		}
	}
	log.Println(best)
	PrintSchedule(p, bestr)

}

func HillClimbOptimize(p []Schedule) {

	var seed []int
	rand.Seed(time.Now().UnixNano())
	for j := 0; j < len(p); j++ {
		seed = append(seed, rand.Intn(10))
		seed = append(seed, rand.Intn(10))
	}

	best := ScheduleCost(p, seed)

	for {

		current := best
		tmp := seed
		for index, t := range seed {
			cost := 0
			//			log.Printf("index : %d, t : %d\n", index, t)
			//			log.Println(seed)
			if t == 0 {
				tmp[index] = t + 1
				cost = ScheduleCost(p, tmp)
				if cost < best {
					best = cost
					seed[index] = tmp[index]
				}
			} else if t == 9 {
				tmp[index] = t - 1
				cost = ScheduleCost(p, tmp)
				if cost < best {
					best = cost
					seed[index] = tmp[index]
				}
			} else {
				tmp[index] = t - 1
				current1 := ScheduleCost(p, tmp)
				tmp[index] = t + 1
				current2 := ScheduleCost(p, tmp)
				if current1 < current2 {
					cost = current1
					tmp[index] = t - 1
				} else {
					cost = current2
					tmp[index] = t + 1
				}
				if cost < best {
					best = cost
					seed[index] = tmp[index]
				}
			}
			//			log.Printf("current cost : %d\n", current)
			//		log.Printf("best cost : %d\n", best)
			//			log.Printf("seed : %v\n", seed)
			//			log.Printf("tmp : %v\n", tmp)
		}

		if best == current {
			//			log.Println("found")
			break
		}

	}

	log.Println(best)
	PrintSchedule(p, seed)

}

func AnnealingOptimize(p []Schedule) {

	T := 100.0
	cool := 0.95
	step := 1

	var seed []int
	rand.Seed(time.Now().UnixNano())
	for j := 0; j < len(p); j++ {
		seed = append(seed, rand.Intn(10))
		seed = append(seed, rand.Intn(10))
	}

	sol := make([]int, len(seed), (cap(seed)+1)*2)
	copy(sol, seed)

	for {

		if T < 0.1 {
			break
		}

//		log.Println(seed)

		i := rand.Intn(12)
		dir := rand.Intn(3) - step

		current := make([]int, len(sol), (cap(sol)+1)*2)
		copy(current, sol)
		current[i] = current[i] + dir

//		log.Println(i)
//		log.Println(current)

		if current[i] < 0 {
			current[i] = seed[i]
		} else if current[i] > 9 {
			current[i] = seed[i]
		}

//		log.Println(current)

		ea := ScheduleCost(p, sol)
		eb := ScheduleCost(p, current)

		p := math.Pow(math.E, -float64(ea+eb)/T)

		log.Println(p)

		if eb < ea {
			sol = current
			copy(sol, current)
		}

		T = T * cool

	}
	log.Println(sol)
	log.Println(ScheduleCost(p, sol))
	PrintSchedule(p, sol)

}
