package polygon

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_KMLWriter(t *testing.T) {
	buf := new(bytes.Buffer)
	sut := NewKMLRespWriter(buf)
	resp := Response{ID: "TEST", Polygon: []Coordinates{{Long: 23.290909090, Lat: -42.2}, {Long: -10.2, Lat: -42.2}}}

	e := sut.Write(resp)
	str := buf.String()

	expected := `<Placemark>
    <name>TEST</name>
    <Polygon>
        <outerBoundaryIs>
            <LinearRing>
                <coordinates> 23.29090909,-42.2 -10.2,-42.2 23.29090909,-42.2 </coordinates>
            </LinearRing>
        </outerBoundaryIs>
    </Polygon>
</Placemark>`
	assert.Nil(t, e)
	assert.Equal(t, expected, str)

}
