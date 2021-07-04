package polygon

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FromRadiusIO_EmptyReq(t *testing.T) {

	io := &ioTest{}
	err := FromRadiusIO(io, io)

	assert.Nil(t, err)
}

func Test_FromRadiusIO_ReturnsOnePolygon(t *testing.T) {

	inputData := []RadiusReq{
		{ID: "TEST",
			Coordinates: Coordinates{Long: float64(0), Lat: float64(0)},
			Radius:      1,
			Edges:       4,
		}}
	io := &ioTest{buffInput: inputData}
	err := FromRadiusIO(io, io)

	assert.Nil(t, err)
	assert.Len(t, io.writes, 1)
	res := io.writes[0]
	assert.Equal(t, "TEST", res.ID)
	assert.Len(t, res.Polygon, 4)
}

func Test_FromRadiusIO_ReturnsThreePolygons(t *testing.T) {

	inputData := []RadiusReq{
		{ID: "TEST_1",
			Coordinates: Coordinates{Long: float64(0), Lat: float64(0)},
			Radius:      1,
			Edges:       4,
		}, {ID: "TEST_2",
			Coordinates: Coordinates{Long: float64(0), Lat: float64(0)},
			Radius:      2,
			Edges:       5,
		}, {ID: "TEST_3",
			Coordinates: Coordinates{Long: float64(1), Lat: float64(1)},
			Radius:      3,
			Edges:       4,
		}}
	io := &ioTest{buffInput: inputData}
	err := FromRadiusIO(io, io)

	assert.Nil(t, err)
	assert.Len(t, io.writes, 3)
	res := io.writes[0]
	assert.Equal(t, "TEST_1", res.ID)
	assert.Len(t, res.Polygon, 4)
	res = io.writes[1]
	assert.Equal(t, "TEST_2", res.ID)
	assert.Len(t, res.Polygon, 5)
	res = io.writes[2]
	assert.Equal(t, "TEST_3", res.ID)
	assert.Len(t, res.Polygon, 4)
}

func Test_FromRadiusIO_WrongLong(t *testing.T) {

	inputData := []RadiusReq{
		{ID: "TEST",
			Coordinates: Coordinates{Long: float64(999), Lat: float64(0)},
			Radius:      1,
			Edges:       4,
		}}
	io := &ioTest{buffInput: inputData}
	err := FromRadiusIO(io, io)

	assert.Nil(t, err)
	assert.Len(t, io.writes, 1)
	res := io.writes[0]
	assert.Equal(t, "TEST", res.ID)
	assert.Len(t, res.Polygon, 0)
	assert.Error(t, res.Error)
}

// TODO TESTs Read and Write Error

type ioTest struct {
	buffInput []RadiusReq
	idx       int
	writes    []PolygonResp
}

func (iot *ioTest) Read() (RadiusReq, error) {
	if len(iot.buffInput) == iot.idx {
		return RadiusReq{}, io.EOF
	}
	res := iot.buffInput[iot.idx]
	iot.idx++
	return res, nil
}

func (iot *ioTest) Write(p PolygonResp) error {
	iot.writes = append(iot.writes, p)
	return nil
}
