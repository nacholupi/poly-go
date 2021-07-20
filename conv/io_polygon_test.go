package conv

import (
	"io"
	"poly-go/polygon"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FromRadiusIO_ReturnsPolygon(t *testing.T) {

	type expResp struct {
		id          string
		coordinates int
	}

	type tCase struct {
		name         string
		req          []request
		expectedResp []expResp
	}

	var testCases = []tCase{
		{
			name:         "Empty Request",
			req:          []request{},
			expectedResp: []expResp{},
		}, {
			name:         "One Polygon",
			req:          []request{nullIslandReq("TEST", 4)},
			expectedResp: []expResp{{id: "TEST", coordinates: 4}},
		}, {
			name: "Three Polygons",
			req: []request{
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
				assert.Equal(t, tc.expectedResp[i].id, res[i].id)
				assert.Len(t, res[i].polygon, tc.expectedResp[i].coordinates)
			}
		})
	}
}

func Test_FromRadiusIO_ReaderFails(t *testing.T) {
	req := []request{nullIslandReq("ID_1", 5), nullIslandReq("ID_2", 4)}
	one := 1
	r := &ioTestReader{input: req, failIdx: &one}
	w := &ioTestWriter{}

	err := FromRadiusIO(r, w)

	assert.Error(t, err)
	assert.Len(t, w.output, 1)
}

func Test_FromRadiusIO_WriterFails(t *testing.T) {
	req := []request{nullIslandReq("ID_1", 5), nullIslandReq("ID_2", 10)}
	one := 1
	r := &ioTestReader{input: req}
	w := &ioTestWriter{failIdx: &one}

	err := FromRadiusIO(r, w)

	assert.Error(t, err)
	assert.Len(t, w.output, 1)
}

func nullIslandReq(id string, edges int) request {
	return request{
		id:          id,
		coordinates: polygon.Coordinates{Long: float64(0), Lat: float64(0)},
		radius:      float64(1),
		edges:       edges,
	}
}

type ioTestReader struct {
	input   []request
	idx     int
	failIdx *int
}

func (r *ioTestReader) read() (request, error) {
	if r.failIdx != nil && r.idx == *r.failIdx {
		return request{}, assert.AnError
	}
	if len(r.input) == r.idx {
		return request{}, io.EOF
	}
	res := r.input[r.idx]
	r.idx++
	return res, nil
}

type ioTestWriter struct {
	output  []response
	idx     int
	failIdx *int
}

func (w *ioTestWriter) write(p response) error {
	if w.failIdx != nil && w.idx == *w.failIdx {
		return assert.AnError
	}
	w.output = append(w.output, p)
	w.idx++
	return nil
}
