package conv

import (
	"io"
)

type csvKml struct {
	reader reader
	writer writer
}

func NewCsvToKmlConverter(r io.Reader, w io.Writer) *csvKml {
	csvReader := NewCSVReqReader(r)
	kmlWriter := NewKMLRespWriter(w)
	return &csvKml{
		reader: csvReader,
		writer: kmlWriter,
	}
}

func (ck *csvKml) CircleToPolygon() error {
	return FromRadiusIO(ck.reader, ck.writer)
}
