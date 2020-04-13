/*Package streamwindow is a helper package, acting as:
* 1. moving 2 slot window for incoming stream data
* 2. filter that ensures only valid points are passing
* 3. enricher, adds data to point that was previously used for filtering and
*    will be needed to perform further calculations
 */
package streamwindow

import (
	"fmt"

	"github.com/fare-estimate/pkg/config"
	"github.com/fare-estimate/pkg/geocalc"
	"github.com/fare-estimate/pkg/models"
)

func enrichPit(p1, p2 *models.PointInTime, speedLimit float64) (*models.PointInTime, error) {
	distMet := geocalc.DistBetween(p1, p2)

	durationSec := p2.Timestamp - p1.Timestamp
	if durationSec <= 0 {
		return nil, fmt.Errorf("%vs seems to be invalid duration", durationSec)
	}

	speedKM := (distMet / float64(durationSec)) * float64(3.6)
	if speedKM > speedLimit {
		return nil, fmt.Errorf("%vkm/h exceeds %vkm/h limit", speedKM, speedLimit)
	}

	p2.SpeedKM = speedKM
	p2.DistanceKM = distMet * 0.001
	p2.DurationH = float64(durationSec / 3600)

	return p2, nil
}

func New(cfg *config.Config) func(pit *models.PointInTime) []*models.PointInTime {
	pitPair := []*models.PointInTime{}

	return func(pit *models.PointInTime) []*models.PointInTime {
		if len(pitPair) == 1 && pitPair[0].ID == pit.ID {
			ePit, err := enrichPit(pitPair[0], pit, cfg.SpeedLimit)
			if err != nil {
				return nil
			}

			pitPair = append(pitPair, ePit)
			return pitPair
		}

		if len(pitPair) == 2 && pitPair[1].ID == pit.ID {
			ePit, err := enrichPit(pitPair[1], pit, cfg.SpeedLimit)
			if err != nil {
				return nil
			}

			pitPair = append(pitPair[1:], ePit)
			return pitPair
		}

		pitPair = nil
		pitPair = append(pitPair, pit)

		return nil
	}
}
