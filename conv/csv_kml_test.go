package conv

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CSV_WrongFields(t *testing.T) {

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
			r := strings.NewReader(tc.csvRecord)
			conv := NewCsvToKmlConverter(r, os.Stdout)

			err := conv.CircleToPolygon()

			assert.Error(t, err)
		})
	}
}
