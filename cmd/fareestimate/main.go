package main

import (
	"log"

	"github.com/fare-estimate/pkg/config"
	"github.com/fare-estimate/pkg/csvwriter"
	"github.com/fare-estimate/pkg/farecalc"
	"github.com/fare-estimate/pkg/filereader"
	"github.com/fare-estimate/pkg/models"
	"github.com/fare-estimate/pkg/streamwindow"
)

func main() {
	cfg := config.New()
	fares := farecalc.New(cfg)
	pairPoints := streamwindow.New(cfg)
	rawPitStream := make(chan []string)

	go func() {
		if err := filereader.ReadFile(cfg.InputFile, rawPitStream); err != nil {
			log.Fatal(err)
		}
	}()

	for rawPIT := range rawPitStream {
		pit, err := models.NewPIT(rawPIT)
		if err != nil {
			log.Fatal(err)
		}

		pitPair := pairPoints(pit)
		if pitPair == nil {
			continue
		}

		fares.AddFare(pitPair[1])
	}

	csvwriter.Write(fares.GetFaresForCSV(), cfg.OutputFile)

	log.Printf("successfully generated %s", cfg.OutputFile)
}
