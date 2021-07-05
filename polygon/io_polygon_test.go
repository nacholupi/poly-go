package polygon

import (
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FromRadiusIO_ReturnsPolygon(t *testing.T) {

	type expResp struct {
		id          string
		coordinates int
		err         error
	}

	type tCase struct {
		name         string
		req          []Request
		expectedResp []expResp
	}

	var testCases = []tCase{
		{
			name:         "Empty Request",
			req:          []Request{},
			expectedResp: []expResp{},
		}, {
			name: "One Polygon",
			req: []Request{
				{ID: "TEST",
					Coordinates: Coordinates{Long: float64(0), Lat: float64(0)},
					Radius:      float64(gradeToKM),
					Edges:       4,
				},
			},
			expectedResp: []expResp{
				{
					id:          "TEST",
					coordinates: 4,
				},
			},
		}, {
			name: "Wrong Long",
			req: []Request{
				{ID: "TEST",
					Coordinates: Coordinates{Long: float64(999), Lat: float64(0)},
					Radius:      float64(gradeToKM),
					Edges:       4,
				},
			},
			expectedResp: []expResp{
				{
					id:          "TEST",
					coordinates: 0,
					err:         fmt.Errorf("Longitude must be greater than -180 and less than 180"),
				},
			},
		}, {
			name: "Three Polygons",
			req: []Request{
				{ID: "TEST_1",
					Coordinates: Coordinates{Long: float64(0), Lat: float64(0)},
					Radius:      float64(gradeToKM),
					Edges:       4,
				}, {ID: "TEST_2",
					Coordinates: Coordinates{Long: float64(1), Lat: float64(1)},
					Radius:      float64(gradeToKM),
					Edges:       5,
				}, {ID: "TEST_3",
					Coordinates: Coordinates{Long: float64(0), Lat: float64(0)},
					Radius:      float64(gradeToKM),
					Edges:       4,
				},
			},
			expectedResp: []expResp{
				{
					id:          "TEST_1",
					coordinates: 4,
				},
				{
					id:          "TEST_2",
					coordinates: 5,
				},
				{
					id:          "TEST_3",
					coordinates: 4,
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
				assert.Equal(t, tc.expectedResp[i].id, res[i].ID)
				assert.Len(t, res[i].Polygon, tc.expectedResp[i].coordinates)
				assert.Equal(t, tc.expectedResp[i].err, res[i].Error)
			}
		})
	}
}

// TODO: TESTs Read and Write Error

type ioTest struct {
	buffInput []Request
	idx       int
	writes    []Response
}

func (iot *ioTest) Read() (Request, error) {
	if len(iot.buffInput) == iot.idx {
		return Request{}, io.EOF
	}
	res := iot.buffInput[iot.idx]
	iot.idx++
	return res, nil
}

func (iot *ioTest) Write(p Response) error {
	iot.writes = append(iot.writes, p)
	return nil
}
