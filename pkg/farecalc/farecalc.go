package farecalc

import (
	"fmt"
	"log"
	"time"

	"github.com/fare-estimate/pkg/config"
	"github.com/fare-estimate/pkg/models"
)

type FareCalc struct {
	cfg   *config.Config
	fares map[uint64]float64
}

func New(cfg *config.Config) *FareCalc {
	return &FareCalc{
		cfg:   cfg,
		fares: map[uint64]float64{},
	}
}

func isNightTariff(ts int64) bool {
	pointTime := time.Unix(ts, 0).UTC()
	midnight := time.Date(pointTime.Year(), pointTime.Month(), pointTime.Day(), 0, 0, 0, 0, time.UTC)
	morning := midnight.Add(5 * time.Hour)

	return pointTime.After(midnight) && pointTime.Before(morning)
}

func (fc *FareCalc) AddFare(pit *models.PointInTime) {
	if _, ok := fc.fares[pit.ID]; !ok {
		fc.fares[pit.ID] = fc.cfg.PickupFee
	}

	tariff := fc.cfg.DayTariff
	if isNightTariff(pit.Timestamp) {
		tariff = fc.cfg.NightTariff
	}

	if pit.SpeedKM < fc.cfg.IdleSpeedKM {
		tariff = pit.DurationH * fc.cfg.IdleTariffPerH
	}

	if pit.DistanceKM == 0 {
		fc.fares[pit.ID] += tariff
	} else {
		fc.fares[pit.ID] += tariff * pit.DistanceKM
	}
}

func (fc *FareCalc) getValidFares() map[uint64]float64 {
	resp := map[uint64]float64{}

	for k, v := range fc.fares {
		if v < fc.cfg.MinFare {
			log.Printf("ID:%v doesn't qualify as valid ride since fare %.2f < minimum %v fare", k, v, fc.cfg.MinFare)
			continue
		}

		resp[k] = v
	}

	return resp
}

func (fc *FareCalc) GetFaresForCSV() [][]string {
	resp := [][]string{{"id_ride", "fare_estimate"}}

	for k, v := range fc.getValidFares() {
		resp = append(resp, []string{
			fmt.Sprintf("%v", k),
			fmt.Sprintf("%.2f", v),
		})
	}

	return resp
}
