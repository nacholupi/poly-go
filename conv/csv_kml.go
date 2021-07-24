package conv

import (
	"io"
)

func NewCsvToKmlConverter(r io.Reader, w io.Writer) *pipe {
	csvReader := newCSVReqReader(r)
	kmlWriter := newKMLRespWriter(w)
	return &pipe{
		in:  csvReader,
		out: kmlWriter,
	}
}
