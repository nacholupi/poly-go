package polygon

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

const gradeToKM = 111.3195 // 1 grade == 111.3195km

func Test_FromRadius_Validations(t *testing.T) {

	type test struct {
		name           string
		coord          Coordinates
		radius         float64
		edges          int
		expectedErrStr string
	}

	var tests = []test{
		{
			name:           "Edges must be greater than 2",
			coord:          Coordinates{Long: float64(0), Lat: float64(0)},
			radius:         float64(1),
			edges:          2,
			expectedErrStr: "Edges must be greater than 2",
		},
		{
			name:           "Radius must be greater than 0: 0",
			coord:          Coordinates{Long: float64(0), Lat: float64(0)},
			radius:         float64(0),
			edges:          3,
			expectedErrStr: "Radius must be greater than 0",
		},
		{
			name:           "Radius must be greater than 0: -1",
			coord:          Coordinates{Long: float64(0), Lat: float64(0)},
			radius:         float64(-1),
			edges:          3,
			expectedErrStr: "Radius must be greater than 0",
		},
		{
			name:           "Latitude must be greater than -90",
			coord:          Coordinates{Long: float64(0), Lat: float64(-90.1)},
			radius:         float64(1),
			edges:          3,
			expectedErrStr: "Latitude must be greater than -90 and less than 90",
		},
		{
			name:           "Latitude must be less than 90",
			coord:          Coordinates{Long: float64(0), Lat: float64(90.1)},
			radius:         float64(1),
			edges:          3,
			expectedErrStr: "Latitude must be greater than -90 and less than 90",
		},
		{
			name:           "Longitude must be greater than -180",
			coord:          Coordinates{Long: float64(-180.1), Lat: float64(0)},
			radius:         float64(1),
			edges:          3,
			expectedErrStr: "Longitude must be greater than -180 and less than 180",
		},
		{
			name:           "Longitude must be less than 180",
			coord:          Coordinates{Long: float64(180.1), Lat: float64(0)},
			radius:         float64(1),
			edges:          3,
			expectedErrStr: "Longitude must be greater than -180 and less than 180",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := FromRadius(tc.coord, tc.radius, tc.edges)

			assert.EqualError(t, err, tc.expectedErrStr)
		})
	}
}

func Test_FromRadius_Results(t *testing.T) {

	type test struct {
		name            string
		coord           Coordinates
		radiusKm        float64
		edges           int
		expectedPoligon []Coordinates
	}

	var tests = []test{
		{
			name:            "Square from Null Island",
			coord:           Coordinates{Long: float64(0), Lat: float64(0)},
			radiusKm:        float64(gradeToKM),
			edges:           4,
			expectedPoligon: []Coordinates{{0, 1}, {1, 0}, {0, -1}, {-1, 0}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			pol, err := FromRadius(tc.coord, tc.radiusKm, tc.edges)
			assert.Nil(t, err)
			assertEqualFloatSlice(t, tc.expectedPoligon, pol, 6)
		})
	}
}

func assertEqualFloatSlice(t *testing.T, expected, actual []Coordinates, precision int) {
	t.Helper()
	assert.Equal(t, len(expected), len(actual))

	for i := range expected {
		e := truncate(expected[i].Lat, precision)
		a := truncate(actual[i].Lat, precision)
		assert.Equal(t, e, a, "wrong latitude in coordinate %d", i)
		e = truncate(expected[i].Long, precision)
		a = truncate(actual[i].Long, precision)
		assert.Equal(t, e, a, "wrong longitude in coordinate %d", i)
	}
}

func truncate(f float64, precision int) float64 {
	p := math.Pow(10, float64(precision))
	return float64(int(f*p)) / p
}
