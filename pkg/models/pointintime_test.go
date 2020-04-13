package models_test

import (
	"testing"

	"github.com/fare-estimate/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewPIT(t *testing.T) {
	cases := map[string]struct {
		input     []string
		assertVal *models.PointInTime
		assertErr func(t *testing.T, err error)
	}{
		"invalid id": {
			input:     []string{"abc", "1.123", "2.123", "123456"},
			assertVal: nil,
			assertErr: func(t *testing.T, err error) {
				require.NotNil(t, err)
			},
		},
		"invalid lat": {
			input:     []string{"1", "abc", "2.123", "123456"},
			assertVal: nil,
			assertErr: func(t *testing.T, err error) {
				require.NotNil(t, err)
			},
		},
		"invalid lng": {
			input:     []string{"1", "1.123", "abc", "123456"},
			assertVal: nil,
			assertErr: func(t *testing.T, err error) {
				require.NotNil(t, err)
			},
		},
		"invalid timestamp": {
			input:     []string{"1", "1.123", "2.123", "abc"},
			assertVal: nil,
			assertErr: func(t *testing.T, err error) {
				require.NotNil(t, err)
			},
		},
		"valid input": {
			input: []string{"1", "1.123", "2.123", "123456"},
			assertVal: &models.PointInTime{
				ID:        1,
				Lat:       1.123,
				Lng:       2.123,
				Timestamp: 123456,
			},
			assertErr: func(t *testing.T, err error) {
				require.Nil(t, err)
			},
		},
	}

	for k, v := range cases {
		t.Run(k, func(t *testing.T) {
			got, err := models.NewPIT(v.input)
			v.assertErr(t, err)
			assert.Equal(t, v.assertVal, got)
		})
	}
}
