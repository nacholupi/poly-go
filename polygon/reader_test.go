package polygon

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Reader_ReturnsRequestThenEOF(t *testing.T) {
	exp := Request{ID: "ID1", Coordinates: Coordinates{Long: 1.1, Lat: 2.2}, Radius: 3.3, Edges: 4}
	csvR := NewCSVReqReader(strings.NewReader("ID1,1.1,2.2,3.3,4"))

	r, err := csvR.Read()
	assert.Equal(t, exp, r)
	assert.Nil(t, err)

	_, err = csvR.Read()
	assert.Equal(t, io.EOF, err)
}

func Test_Reader_TrimLeadingSpace(t *testing.T) {
	exp := Request{ID: "ID1", Coordinates: Coordinates{Long: 1.1, Lat: 2.2}, Radius: 3.3, Edges: 4}
	csvR := NewCSVReqReader(strings.NewReader("ID1,   1.1, 2.2, 3.3, 4"))

	r, _ := csvR.Read()

	assert.Equal(t, exp, r)
}
