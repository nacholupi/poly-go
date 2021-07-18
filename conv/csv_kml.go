package conv

import (
	"io"
	"poly-go/polygon"
)

type csvKml struct {
	reader polygon.Reader
}

func NewCsvToKmlConverter(r io.Reader, w io.Writer) *csvKml {
	csvReader := polygon.NewCSVReqReader(r)
	return &csvKml{
		reader: csvReader,
	}
}

func (ck *csvKml) CircleToPolygon() error {
	return polygon.FromRadiusIO(ck.reader, nil)
}
