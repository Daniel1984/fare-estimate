package models

import "strconv"

type PointInTime struct {
	ID         uint64
	Lat        float64
	Lng        float64
	Timestamp  int64
	SpeedKM    float64
	DistanceKM float64
	DurationH  float64
	err        error
}

func (pit *PointInTime) withID(id string) {
	if pit.err != nil {
		return
	}

	pit.ID, pit.err = strconv.ParseUint(id, 10, 64)
}

func (pit *PointInTime) withLAT(lat string) {
	if pit.err != nil {
		return
	}

	pit.Lat, pit.err = strconv.ParseFloat(lat, 64)
}

func (pit *PointInTime) withLNG(lng string) {
	if pit.err != nil {
		return
	}

	pit.Lng, pit.err = strconv.ParseFloat(lng, 64)
}

func (pit *PointInTime) withTimestamp(ts string) {
	if pit.err != nil {
		return
	}

	pit.Timestamp, pit.err = strconv.ParseInt(ts, 10, 64)
}

func NewPIT(cols []string) (*PointInTime, error) {
	pit := &PointInTime{}
	pit.withID(cols[0])
	pit.withLAT(cols[1])
	pit.withLNG(cols[2])
	pit.withTimestamp(cols[3])

	if pit.err != nil {
		return nil, pit.err
	}

	return pit, nil
}
