package polygon

import (
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FromRadiusIO_ReturnsPolygon(t *testing.T) {

	type tCase struct {
		name         string
		req          []RadiusReq
		expectedResp []PolygonResp
	}

	var testCases = []tCase{
		{
			name:         "Empty Request",
			req:          []RadiusReq{},
			expectedResp: []PolygonResp{},
		}, {
			name: "One Polygon",
			req: []RadiusReq{
				{ID: "TEST",
					Coordinates: Coordinates{Long: float64(0), Lat: float64(0)},
					Radius:      float64(gradeToKM),
					Edges:       4,
				},
			},
			expectedResp: []PolygonResp{
				{
					ID:      "TEST",
					Polygon: []Coordinates{{0, 1}, {1, 0}, {0, -1}, {-1, 0}},
				},
			},
		}, {
			name: "Wrong Long",
			req: []RadiusReq{
				{ID: "TEST",
					Coordinates: Coordinates{Long: float64(999), Lat: float64(0)},
					Radius:      float64(gradeToKM),
					Edges:       4,
				},
			},
			expectedResp: []PolygonResp{
				{
					ID:      "TEST",
					Polygon: []Coordinates{},
					Error:   fmt.Errorf("Longitude must be greater than -180 and less than 180"),
				},
			},
		}, {
			name: "Three Polygons",
			req: []RadiusReq{
				{ID: "TEST_1",
					Coordinates: Coordinates{Long: float64(0), Lat: float64(0)},
					Radius:      float64(gradeToKM),
					Edges:       4,
				}, {ID: "TEST_2",
					Coordinates: Coordinates{Long: float64(0), Lat: float64(0)},
					Radius:      float64(gradeToKM),
					Edges:       4,
				}, {ID: "TEST_3",
					Coordinates: Coordinates{Long: float64(0), Lat: float64(0)},
					Radius:      float64(gradeToKM),
					Edges:       4,
				},
			},
			expectedResp: []PolygonResp{
				{
					ID:      "TEST_1",
					Polygon: []Coordinates{{0, 1}, {1, 0}, {0, -1}, {-1, 0}},
				}, {
					ID:      "TEST_2",
					Polygon: []Coordinates{{0, 1}, {1, 0}, {0, -1}, {-1, 0}},
				}, {
					ID:      "TEST_3",
					Polygon: []Coordinates{{0, 1}, {1, 0}, {0, -1}, {-1, 0}},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			io := &ioTest{buffInput: tc.req}
			err := FromRadiusIO(io, io)

			res := io.writes
			assert.Nil(t, err)
			assert.Len(t, res, len(tc.expectedResp))
			for i := range res {
				assert.Equal(t, tc.expectedResp[i].ID, res[i].ID)
				assertEqualFloatSlice(t, tc.expectedResp[i].Polygon, res[i].Polygon, 6)
				assert.Equal(t, tc.expectedResp[i].Error, res[i].Error)
			}
		})
	}
}

// TODO TESTs Read and Write Error

type ioTest struct {
	buffInput []RadiusReq
	idx       int
	writes    []PolygonResp
}

func (iot *ioTest) Read() (RadiusReq, error) {
	if len(iot.buffInput) == iot.idx {
		return RadiusReq{}, io.EOF
	}
	res := iot.buffInput[iot.idx]
	iot.idx++
	return res, nil
}

func (iot *ioTest) Write(p PolygonResp) error {
	iot.writes = append(iot.writes, p)
	return nil
}
