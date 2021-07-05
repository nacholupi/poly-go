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
			name:         "One Polygon",
			req:          []Request{nullIslandReq("TEST", 4)},
			expectedResp: []expResp{{id: "TEST", coordinates: 4}},
		}, {
			name: "Wrong Long",
			req: []Request{
				{ID: "TEST",
					Coordinates: Coordinates{Long: float64(999), Lat: float64(0)},
					Radius:      1,
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
				nullIslandReq("ID_1", 4),
				nullIslandReq("ID_2", 5),
				nullIslandReq("ID_3", 10)},
			expectedResp: []expResp{
				{id: "ID_1", coordinates: 4},
				{id: "ID_2", coordinates: 5},
				{id: "ID_3", coordinates: 10}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := &ioTestReader{input: tc.req}
			w := &ioTestWriter{}

			err := FromRadiusIO(r, w)

			res := w.output
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

func Test_FromRadiusIO_ReaderFails(t *testing.T) {
	req := []Request{nullIslandReq("ID_1", 5), nullIslandReq("ID_2", 4)}
	one := 1
	r := &ioTestReader{input: req, failIdx: &one}
	w := &ioTestWriter{}

	err := FromRadiusIO(r, w)

	assert.Error(t, err)
	assert.Len(t, w.output, 1)
}

func Test_FromRadiusIO_WriterFails(t *testing.T) {
	req := []Request{nullIslandReq("ID_1", 5), nullIslandReq("ID_2", 10)}
	one := 1
	r := &ioTestReader{input: req}
	w := &ioTestWriter{failIdx: &one}

	err := FromRadiusIO(r, w)

	assert.Error(t, err)
	assert.Len(t, w.output, 1)
}

func nullIslandReq(id string, edges int) Request {
	return Request{ID: id,
		Coordinates: Coordinates{Long: float64(0), Lat: float64(0)},
		Radius:      float64(1),
		Edges:       edges,
	}
}

type ioTestReader struct {
	input   []Request
	idx     int
	failIdx *int
}

func (r *ioTestReader) Read() (Request, error) {
	if r.failIdx != nil && r.idx == *r.failIdx {
		return Request{}, assert.AnError
	}
	if len(r.input) == r.idx {
		return Request{}, io.EOF
	}
	res := r.input[r.idx]
	r.idx++
	return res, nil
}

type ioTestWriter struct {
	output  []Response
	idx     int
	failIdx *int
}

func (w *ioTestWriter) Write(p Response) error {
	if w.failIdx != nil && w.idx == *w.failIdx {
		return assert.AnError
	}
	w.output = append(w.output, p)
	w.idx++
	return nil
}
