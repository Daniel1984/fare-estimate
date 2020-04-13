package streamwindow_test

import (
	"testing"

	"github.com/fare-estimate/pkg/config"
	"github.com/fare-estimate/pkg/models"
	"github.com/fare-estimate/pkg/streamwindow"
	"github.com/stretchr/testify/assert"
)

func TestNewStreamWindow(t *testing.T) {
	cases := map[string]struct {
		input     []*models.PointInTime
		assertVal []*models.PointInTime
	}{
		"single point": {
			input:     []*models.PointInTime{&models.PointInTime{}},
			assertVal: nil,
		},
		"2 points invalid time": {
			input: []*models.PointInTime{
				&models.PointInTime{ID: 1, Timestamp: 100},
				&models.PointInTime{ID: 1, Timestamp: 100},
			},
			assertVal: nil,
		},
		"2 points invalid speed": {
			input: []*models.PointInTime{
				&models.PointInTime{
					ID:        1,
					Lat:       37.966660,
					Lng:       23.728308,
					Timestamp: 50,
				},
				&models.PointInTime{
					ID:        1,
					Lat:       38.966660,
					Lng:       23.128308,
					Timestamp: 100,
				},
			},
			assertVal: nil,
		},
		"2 valid points different id's": {
			input: []*models.PointInTime{
				&models.PointInTime{
					ID:        1,
					Lat:       37.966660,
					Lng:       23.728308,
					Timestamp: 1405594957,
				},
				&models.PointInTime{
					ID:        2,
					Lat:       37.966627,
					Lng:       23.728263,
					Timestamp: 1405594966,
				},
			},
			assertVal: nil,
		},
		"2 valid points": {
			input: []*models.PointInTime{
				&models.PointInTime{
					ID:        1,
					Lat:       37.966660,
					Lng:       23.728308,
					Timestamp: 1405594957,
				},
				&models.PointInTime{
					ID:        1,
					Lat:       37.966627,
					Lng:       23.728263,
					Timestamp: 1405594966,
				},
			},
			assertVal: []*models.PointInTime{
				&models.PointInTime{
					ID:        1,
					Lat:       37.966660,
					Lng:       23.728308,
					Timestamp: 1405594957,
				},
				&models.PointInTime{
					ID:         1,
					Lat:        37.966627,
					Lng:        23.728263,
					Timestamp:  1405594966,
					SpeedKM:    2.1550435801161765,
					DistanceKM: 0.005387608950290441,
				},
			},
		},
	}

	for k, v := range cases {
		t.Run(k, func(t *testing.T) {
			mw := streamwindow.New(&config.Config{SpeedLimit: 100})
			var got []*models.PointInTime

			for _, poi := range v.input {
				got = mw(poi)
			}

			assert.Equal(t, v.assertVal, got)
		})
	}
}
