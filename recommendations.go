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
	util.LoadMovieLens()

}

func main() {
	//	TestdefaultDataset()

	TestMovieLens()

}
