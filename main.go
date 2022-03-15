package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Hero struct {
	name        string
	proficiency string
}

func main() {
	f, err := os.Open("heroes.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	allHeroes := createDotaList(data)

	fmt.Printf("%+v\n", allHeroes)
}

func createDotaList(data [][]string) []Hero {
	var allHeroes []Hero
	for i, line := range data {
		if i > 0 { //skip line 1
			var record Hero
			for j, field := range line {
				if j == 0 {
					record.name = field
				} else if j == 1 {
					record.proficiency = field
				}
			}
			allHeroes = append(allHeroes, record)
		}
	}
	return allHeroes
}
