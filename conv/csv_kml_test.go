package conv

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CircleToPolygon_Parsing_Error(t *testing.T) {

	type tCase struct {
		name string
		csv  string
	}
	var tCases = []tCase{
		{
			name: "Wrong number of fields",
			csv:  "ID1,123,123",
		},
		{
			name: "Wrong Long",
			csv:  "ID1,wrong,1,2,3",
		},
		{
			name: "Wrong Lat (string)",
			csv:  "ID1,1,wrong,2,3",
		},
		{
			name: "Wrong Lat (greater than 90) ",
			csv:  "ID1,1,190,2,3",
		},
		{
			name: "Wrong Radius",
			csv:  "ID1,1,2,wrong,3",
		},
		{
			name: "Wrong Edge (string)",
			csv:  "ID1,1,2,3,wrong",
		},
		{
			name: "Wrong Edge (float)",
			csv:  "ID1,1,2,3,4.1",
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(t *testing.T) {
			r := strings.NewReader(tc.csv)
			conv := NewCsvToKmlConverter(r, os.Stdout)

			err := conv.CircleToPolygon()

			assert.Error(t, err)
		})
	}
}

func Test_CircleToPolygon_When_Reader_Returns_Error(t *testing.T) {
	r := new(errReader)
	conv := NewCsvToKmlConverter(r, os.Stdout)

	err := conv.CircleToPolygon()

	assert.Error(t, err)
}

func Test_CircleToPolygon_KML_Responses(t *testing.T) {

	type tCase struct {
		name     string
		csv      string
		expected string
		err      bool
	}
	var tCases = []tCase{
		{
			name: "CSV with single row - 3 edges",
			csv:  "ID1,1,-89,2,3",
			expected: "  <Placemark>\n" +
				"    <name>ID1</name>\n" +
				"    <Polygon>\n" +
				"      <outerBoundaryIs>\n" +
				"        <LinearRing>\n" +
				"          <coordinates> 1.00000000,-88.98203369 1.89953277,-89.00886103 0.10046723,-89.00886103" +
				" 1.00000000,-88.98203369</coordinates>\n" +
				"        </LinearRing>\n" +
				"      </outerBoundaryIs>\n" +
				"    </Polygon>\n" +
				"  </Placemark>",
			err: false,
		},
		{
			name: "CSV with single row - trimmed long, lat",
			csv:  "ID1,       1,  -89,2,3",
			expected: "  <Placemark>\n" +
				"    <name>ID1</name>\n" +
				"    <Polygon>\n" +
				"      <outerBoundaryIs>\n" +
				"        <LinearRing>\n" +
				"          <coordinates> 1.00000000,-88.98203369 1.89953277,-89.00886103 0.10046723,-89.00886103" +
				" 1.00000000,-88.98203369</coordinates>\n" +
				"        </LinearRing>\n" +
				"      </outerBoundaryIs>\n" +
				"    </Polygon>\n" +
				"  </Placemark>",
			err: false,
		},
		{
			name: "CSV with two rows - 3 and 10 edges",
			csv: "ID1,1,-89,2,3\n" +
				"ID2,0,0,1,4",
			expected: "  <Placemark>\n" +
				"    <name>ID1</name>\n" +
				"    <Polygon>\n" +
				"      <outerBoundaryIs>\n" +
				"        <LinearRing>\n" +
				"          <coordinates> 1.00000000,-88.98203369 1.89953277,-89.00886103 0.10046723,-89.00886103" +
				" 1.00000000,-88.98203369</coordinates>\n" +
				"        </LinearRing>\n" +
				"      </outerBoundaryIs>\n" +
				"    </Polygon>\n" +
				"  </Placemark>\n" +
				"  <Placemark>\n" +
				"    <name>ID2</name>\n" +
				"    <Polygon>\n" +
				"      <outerBoundaryIs>\n" +
				"        <LinearRing>\n" +
				"          <coordinates> 0.00000000,0.00898315 0.00898315,0.00000000 0.00000000,-0.00898315 " +
				"-0.00898315,-0.00000000 0.00000000,0.00898315</coordinates>\n" +
				"        </LinearRing>\n" +
				"      </outerBoundaryIs>\n" +
				"    </Polygon>\n" +
				"  </Placemark>",
			err: false,
		},
		{
			name: "CSV with two rows - the second fails",
			csv: "ID1,1,-89,2,3\n" +
				"ID2,0,1000,1,4",
			expected: "  <Placemark>\n" +
				"    <name>ID1</name>\n" +
				"    <Polygon>\n" +
				"      <outerBoundaryIs>\n" +
				"        <LinearRing>\n" +
				"          <coordinates> 1.00000000,-88.98203369 1.89953277,-89.00886103 0.10046723,-89.00886103" +
				" 1.00000000,-88.98203369</coordinates>\n" +
				"        </LinearRing>\n" +
				"      </outerBoundaryIs>\n" +
				"    </Polygon>\n" +
				"  </Placemark>",
			err: true,
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(t *testing.T) {
			r := strings.NewReader(tc.csv)
			w := new(bytes.Buffer)
			conv := NewCsvToKmlConverter(r, w)

			err := conv.CircleToPolygon()

			assert.Equal(t, tc.err, err != nil)
			assert.Equal(t, tc.expected, w.String())
		})
	}
}

type errReader struct{}

func (*errReader) Read(b []byte) (n int, err error) { return 0, fmt.Errorf("error") }
