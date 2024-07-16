package main

import (
	"cannabits/strainparser"
	"fmt"
	"time"
)

// minumum # of similarities to include a strain in the ranking
const THRESHOLD = 2

// specify path to csv file
const DATAPATH = "/Users/blaze/Development/experiments/cannaparser/data/cannabis.csv"

// build a query with desired flavors and effects
var desiredEffects = []string{"Creative", "Focused"}
var desiredFlavors = []string{"Strawberry", "Mango"}
var desiredType = "sativa"

func main() {
	fmt.Println("== Cannaparser v0.0.1 ==\n[x] Loading CSV...")
	report := strainparser.ParseStrains(DATAPATH)
	fmt.Printf("[x] Parsed %d strains in %d milliseconds\n", len(report.Strains), report.IgnoredCount)

	start := time.Now().UnixMicro()
	var encodings []*strainparser.StrainBitEncoding
	for _, s := range report.Strains {
		encodings = append(encodings, strainparser.GetStrainEncodings(s))
	}
	fmt.Printf("[x] Encoded strains in %d microseconds\n", time.Now().UnixMicro()-start)
	start = time.Now().UnixMicro()

	query := strainparser.NewQueryEncoding(desiredType, desiredEffects, desiredFlavors)

	ranker := strainparser.NewStrainsHeap()
	for idx, enc := range encodings {
		sim := query.GetOneSimilarityIgnoreType(enc)
		//to ignore type:
		//  query.GetOneSimilarityIgnoreType(enc)
		//just flavors:
		//  query.GetFlavorsOneSimilarity(enc)
		//just effects:
		//	query.GetFlavorsOneSimilarity(enc)

		if sim < THRESHOLD {
			continue
		}

		//weight the similarity score slightly to prefer higher Rated items
		weightAmount := int(report.Strains[idx].Rating * 0.2)
		ranker.Insert(&report.Strains[idx], sim+weightAmount)
	}

	var top10 []*strainparser.StrainEntry
	for i := 0; i < 10; i += 1 {
		//prevent underflow
		if ranker.Size == 0 {
			break
		}
		top10 = append(top10, ranker.Extract())
	}
	fmt.Printf("[x] Got best results in %d microseconds\n\n ==== RESULTS ====", time.Now().UnixMicro()-start)
	for i := 0; i < 10; i++ {
		fmt.Printf("\n(%d) %v - %v \n[Rating]: %f\n[Effects]: %v\n[Flavors]: %v\n", i+1, top10[i].Name, top10[i].Type, top10[i].Rating, top10[i].Effects, top10[i].Flavors)

	}

}
