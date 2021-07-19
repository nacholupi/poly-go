package conv

import (
	"io"
	"poly-go/polygon"
)

type csvKml struct {
	reader polygon.Reader
	writer polygon.Writer
}

func NewCsvToKmlConverter(r io.Reader, w io.Writer) *csvKml {
	csvReader := polygon.NewCSVReqReader(r)
	kmlWriter := polygon.NewKMLRespWriter(w)
	return &csvKml{
		reader: csvReader,
		writer: kmlWriter,
	}
}

func (ck *csvKml) CircleToPolygon() error {
	return polygon.FromRadiusIO(ck.reader, ck.writer)
}
