package farecalc_test

import (
	"testing"

	"github.com/fare-estimate/pkg/config"
	"github.com/fare-estimate/pkg/farecalc"
	"github.com/fare-estimate/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestGetFaresForCSV(t *testing.T) {
	cases := map[string]struct {
		input     []*models.PointInTime
		assertVal [][]string
	}{
		"night ride": {
			input: []*models.PointInTime{
				&models.PointInTime{
					ID:         1,
					SpeedKM:    50,
					DurationH:  0.1,
					DistanceKM: 0.1,
					Timestamp:  1405562400,
				},
				&models.PointInTime{
					ID:         1,
					SpeedKM:    50,
					DurationH:  0.2,
					DistanceKM: 0.2,
					Timestamp:  1405562400,
				},
			},
			assertVal: [][]string{
				{"id_ride", "fare_estimate"},
				{"1", "2.20"},
			},
		},
		"day ride": {
			input: []*models.PointInTime{
				&models.PointInTime{
					ID:         1,
					SpeedKM:    50,
					DurationH:  0.1,
					DistanceKM: 0.1,
					Timestamp:  1405537200,
				},
				&models.PointInTime{
					ID:         1,
					SpeedKM:    50,
					DurationH:  0.2,
					DistanceKM: 0.2,
					Timestamp:  1405537200,
				},
			},
			assertVal: [][]string{
				{"id_ride", "fare_estimate"},
				{"1", "1.60"},
			},
		},
		"night ride IDLE": {
			input: []*models.PointInTime{
				&models.PointInTime{
					ID:         1,
					SpeedKM:    5,
					DurationH:  0.1,
					DistanceKM: 0.1,
					Timestamp:  1405562400,
				},
				&models.PointInTime{
					ID:         1,
					SpeedKM:    5,
					DurationH:  0.2,
					DistanceKM: 0.2,
					Timestamp:  1405562400,
				},
			},
			assertVal: [][]string{
				{"id_ride", "fare_estimate"},
				{"1", "1.25"},
			},
		},
		"day ride IDLE": {
			input: []*models.PointInTime{
				&models.PointInTime{
					ID:         1,
					SpeedKM:    5,
					DurationH:  0.1,
					DistanceKM: 0.1,
					Timestamp:  1405537200,
				},
				&models.PointInTime{
					ID:         1,
					SpeedKM:    5,
					DurationH:  0.2,
					DistanceKM: 0.2,
					Timestamp:  1405537200,
				},
			},
			assertVal: [][]string{
				{"id_ride", "fare_estimate"},
				{"1", "1.25"},
			},
		},
		"day ride IDLE 0 distance": {
			input: []*models.PointInTime{
				&models.PointInTime{
					ID:         1,
					SpeedKM:    5,
					DurationH:  0.1,
					DistanceKM: 0,
					Timestamp:  1405537200,
				},
				&models.PointInTime{
					ID:         1,
					SpeedKM:    5,
					DurationH:  0.2,
					DistanceKM: 0,
					Timestamp:  1405537200,
				},
			},
			assertVal: [][]string{
				{"id_ride", "fare_estimate"},
				{"1", "2.50"},
			},
		},
	}

	for k, v := range cases {
		t.Run(k, func(t *testing.T) {
			fc := farecalc.New(&config.Config{
				PickupFee:      1,
				DayTariff:      2,
				NightTariff:    4,
				IdleSpeedKM:    10,
				MinFare:        1,
				IdleTariffPerH: 5,
			})
			for _, pit := range v.input {
				fc.AddFare(pit)
			}

			got := fc.GetFaresForCSV()
			assert.Equal(t, v.assertVal, got)
		})
	}
}
