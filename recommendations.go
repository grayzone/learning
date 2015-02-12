package main

import (
	"github.com/grayzone/learning/util"
	"log"
)

func TestdefaultDataset() {
	d := util.Initialize()

	//	log.Println(sim_distance(d, "test", "Jack Matthews"))
	log.Println(util.Sim_distance(d, "Lisa Rose", "Gene Seymour"))
	log.Println(util.Sim_pearson(d, "Lisa Rose", "Gene Seymour"))
	log.Println(util.Sim_pearson(d, "Toby", "Claudia Puig"))

	log.Println(util.TopMatches(d, "Toby", 3, util.Sim_pearson))

	log.Println(util.GetRecommendations(d, "Toby", util.Sim_pearson))
	log.Println(util.GetRecommendations(d, "Toby", util.Sim_distance))
	log.Println(util.GetRecommendations(d, "Lisa Rose", util.Sim_distance))

	d2 := util.TransformPrefs(d)
	log.Println(util.TopMatches(d2, "Superman Returns", 5, util.Sim_pearson))
	log.Println(util.GetRecommendations(d2, "Just My Luck", util.Sim_pearson))

	itemSim := util.CalculateSimilarItems(d, 10)

	util.GetRecommendationItems(d, itemSim, "Toby")
}

func TestMovieLens() {

	title := util.LoadMovieTitles()
	d := util.LoadMovieData(title)
	//	log.Println(d["87"])
	//	log.Println(util.GetRecommendations(d, "87", util.Sim_pearson)[:30])

	itemsim := util.CalculateSimilarItems(d, 50)

	log.Println(util.GetRecommendationItems(d, itemsim, "87")[:30])

}

func TestOptimization() {
	util.InitOpt()
	util.LoadScheduleData()
}

func main() {
	//	TestdefaultDataset()

	// TestMovieLens()
	TestOptimization()

}
