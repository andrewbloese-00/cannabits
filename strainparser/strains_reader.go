package strainparser

import (
	"encoding/csv"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type StrainEntry struct {
	Name        string
	Type        string
	Rating      float64
	Effects     []string
	Flavors     []string
	Description string
}

type StrainParserReport struct {
	Strains      []StrainEntry
	Flavors      []string
	Effects      []string
	IgnoredCount int
	ParseTimeMS  int
}

// parses strains from csv at DATAPATH and returns a report
func ParseStrains(datapath string) StrainParserReport {
	var ignoreCount int = 0
	start := time.Now().UnixMilli()
	records, err := readStrainsCSV(datapath)

	if err != nil {
		panic(err)
	}
	var strains []StrainEntry

	// 'sets'
	uniq_effects, uniq_flavors := make(map[string]bool), make(map[string]bool)
	for _, record := range records {
		ratingFloat, convErr := strconv.ParseFloat(record[2], 64)
		if convErr != nil || ratingFloat == 0.0 {
			// fmt.Println("Failed to parse rating as non-zero float, skipping entry.")
			ignoreCount += 1
			continue
		}

		// if no flavors or effects recorded ignore
		if record[3] == "None" || record[4] == "None" {
			// fmt.Println("Strain had 'None' as flavor or effect, skipping due to insufficient data")
			ignoreCount += 1
			continue
		}

		// parse out flavors and effects, also tracking all unique
		effects, flavors := strings.Split(record[3], ","), strings.Split(record[4], ",")

		for _, e := range effects {
			if _, ok := uniq_effects[e]; !ok {
				uniq_effects[e] = true
			}
		}
		for _, f := range flavors {
			if _, ok := uniq_flavors[f]; !ok {
				uniq_flavors[f] = true
			}
		}

		strains = append(strains, StrainEntry{
			Name:        record[0],
			Type:        record[1],
			Rating:      ratingFloat,
			Effects:     effects,
			Flavors:     flavors,
			Description: record[5],
		})
	}

	return StrainParserReport{
		Strains:      strains,
		Effects:      mapToSet(uniq_effects),
		Flavors:      mapToSet(uniq_flavors),
		IgnoredCount: ignoreCount,
		ParseTimeMS:  int((time.Now().UnixMilli()) - start),
	}
}

// helper - convert map into an alphabetic orderered list
func mapToSet(inputMap map[string]bool) []string {
	var set []string
	for k := range inputMap {
		if k != "Dry" && k != "Mouth" {
			set = append(set, k)
		}
	}
	sort.Strings(set)
	return set
}

// err response 'records'
var default_records = [][]string{}

// helper - read rows as records
func readStrainsCSV(datapath string) ([][]string, error) {
	file, err := os.Open(datapath)
	if err != nil {
		return default_records, err
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return default_records, err
	}
	return records, nil
}
