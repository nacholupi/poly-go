package conv

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_KMLWriteHeader(t *testing.T) {
	buf := new(bytes.Buffer)
	sut := NewKMLRespWriter(buf)

	e := sut.writeHeader()

	str := buf.String()
	expected := `<?xml version="1.0" encoding="UTF-8"?>` + "\n" +
		`<kml xmlns="http://www.opengis.net/kml/2.2">` + "\n" +
		`  <Document>` + "\n"

	assert.Nil(t, e)
	assert.Equal(t, expected, str)

}

func Test_KMLWriteFooter(t *testing.T) {
	buf := new(bytes.Buffer)
	sut := NewKMLRespWriter(buf)

	e := sut.writeFooter()

	str := buf.String()
	expected := `  </Document>` + "\n" +
		`</kml>`

	assert.Nil(t, e)
	assert.Equal(t, expected, str)

}
