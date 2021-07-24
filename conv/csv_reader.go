package conv

import (
	"encoding/csv"
	"io"
	"poly-go/polygon"
	"strconv"

	"github.com/pkg/errors"
)

type csvReqReader struct {
	csvReader *csv.Reader
}

// CSV Format row: id, long, lat, radius, edges
func newCSVReqReader(r io.Reader) *csvReqReader {
	csvr := csv.NewReader(r)
	csvr.FieldsPerRecord = 5
	csvr.TrimLeadingSpace = true
	return &csvReqReader{csvReader: csvr}
}

func (r *csvReqReader) read() (request, error) {
	rec, err := r.csvReader.Read()
	if err != nil {
		return request{}, err
	}
	req, err := r.parseRecord(rec)

	if err != nil {
		return req, errors.Wrap(err, "Parsing CSV error")
	}

	return req, err
}

func (r *csvReqReader) parseRecord(record []string) (request, error) {
	long, err := strconv.ParseFloat(record[1], 64)
	if err != nil {
		return request{}, err
	}

	lat, err := strconv.ParseFloat(record[2], 64)
	if err != nil {
		return request{}, err
	}

	rad, err := strconv.ParseFloat(record[3], 64)
	if err != nil {
		return request{}, err
	}

	edges, err := strconv.ParseInt(record[4], 10, 32)
	if err != nil {
		return request{}, err
	}

	req := request{
		id:          record[0],
		coordinates: polygon.Coordinates{Long: long, Lat: lat},
		radius:      rad,
		edges:       int(edges),
	}

	return req, nil
}
