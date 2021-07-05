package polygon

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Reader_WrongFields(t *testing.T) {

	type tCase struct {
		name      string
		csvRecord string
	}
	var tCases = []tCase{
		{
			name:      "Wrong number of fields",
			csvRecord: "ID1,123,123",
		},
		{
			name:      "Wrong Long",
			csvRecord: "ID1,wrong,1,2,3",
		},
		{
			name:      "Wrong Lat",
			csvRecord: "ID1,1,wrong,2,3",
		},
		{
			name:      "Wrong Radius",
			csvRecord: "ID1,1,2,wrong,3",
		},
		{
			name:      "Wrong Edge (string)",
			csvRecord: "ID1,1,2,3,wrong",
		},
		{
			name:      "Wrong Edge (float)",
			csvRecord: "ID1,1,2,3,4.1",
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(t *testing.T) {
			csvR := NewCSVReqReader(strings.NewReader(tc.csvRecord))

			_, err := csvR.Read()

			assert.Error(t, err)
		})
	}
}

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
