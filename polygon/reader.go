package polygon

import (
	"encoding/csv"
	"io"
	"strconv"

	"github.com/pkg/errors"
)

type csvReqReader struct {
	csvReader *csv.Reader
}

// CSV Format row: id, long, lat, radius, edges
func NewCSVReqReader(r io.Reader) *csvReqReader {
	csvr := csv.NewReader(r)
	csvr.FieldsPerRecord = 5
	csvr.TrimLeadingSpace = true
	return &csvReqReader{csvReader: csvr}
}

func (r *csvReqReader) Read() (Request, error) {
	rec, err := r.csvReader.Read()
	if err != nil {
		return Request{}, err
	}
	req, err := r.parseRecord(rec)

	if err != nil {
		return req, errors.Wrap(err, "Parsing CSV error")
	}

	return req, err
}

func (r *csvReqReader) parseRecord(record []string) (Request, error) {
	long, err := strconv.ParseFloat(record[1], 64)
	if err != nil {
		return Request{}, err
	}

	lat, err := strconv.ParseFloat(record[2], 64)
	if err != nil {
		return Request{}, err
	}

	rad, err := strconv.ParseFloat(record[3], 64)
	if err != nil {
		return Request{}, err
	}

	edges, err := strconv.ParseInt(record[4], 10, 32)
	if err != nil {
		return Request{}, err
	}

	req := Request{
		ID:          record[0],
		Coordinates: Coordinates{Long: long, Lat: lat},
		Radius:      rad,
		Edges:       int(edges),
	}

	return req, nil
}
