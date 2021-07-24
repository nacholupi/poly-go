package conv

import (
	"io"
	"poly-go/polygon"
	"testing"

	"github.com/stretchr/testify/assert"
)

//TODO: remove
func Test_FromRadiusIO_ReaderFails(t *testing.T) {
	req := []request{nullIslandReq("ID_1", 5), nullIslandReq("ID_2", 4)}
	one := 1
	r := &ioTestReader{input: req, failIdx: &one}
	w := &ioTestWriter{}
	p := pipe{in: r, out: w}

	err := p.CircleToPolygon()

	assert.Error(t, err)
	assert.Len(t, w.output, 1)
}

//TODO: remove
func Test_FromRadiusIO_WriterFails(t *testing.T) {
	req := []request{nullIslandReq("ID_1", 5), nullIslandReq("ID_2", 10)}
	one := 1
	r := &ioTestReader{input: req}
	w := &ioTestWriter{failIdx: &one}
	p := pipe{in: r, out: w}

	err := p.CircleToPolygon()

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

func (w *ioTestWriter) writeHeader() error { return nil }

func (w *ioTestWriter) writeFooter() error { return nil }
