package polygon

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_KMLWrite(t *testing.T) {
	buf := new(bytes.Buffer)
	sut := NewKMLRespWriter(buf)
	resp := Response{ID: "TEST", Polygon: []Coordinates{{Long: 23.290909090, Lat: -42.2}, {Long: -10.2, Lat: -42.2}}}

	e := sut.Write(resp)
	str := buf.String()

	expected := "  <Placemark>\n" +
		"    <name>TEST</name>\n" +
		"    <Polygon>\n" +
		"      <outerBoundaryIs>\n" +
		"        <LinearRing>\n" +
		"          <coordinates> 23.29090909,-42.20000000 -10.20000000,-42.20000000 23.29090909,-42.20000000</coordinates>\n" +
		"        </LinearRing>\n" +
		"      </outerBoundaryIs>\n" +
		"    </Polygon>\n" +
		"  </Placemark>"
	assert.Nil(t, e)
	assert.Equal(t, expected, str)

}

func Test_KMLWriteHeader(t *testing.T) {
	buf := new(bytes.Buffer)
	sut := NewKMLRespWriter(buf)

	e := sut.WriteHeader()

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

	e := sut.WriteFooter()

	str := buf.String()
	expected := `  </Document>` + "\n" +
		`</kml>`

	assert.Nil(t, e)
	assert.Equal(t, expected, str)

}
