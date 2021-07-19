package polygon

import (
	"fmt"
	"io"
	"math"
)

type Coordinates struct {
	Long float64
	Lat  float64
}

const earthRadiusKM = float64(6378.137)

func FromRadius(coord Coordinates, radiusKm float64, edges int) ([]Coordinates, error) {

	err := validateInput(coord, radiusKm, edges)
	if err != nil {
		return nil, err
	}

	pol := polygonFromRadius(coord, radiusKm, edges)

	return pol, nil
}

func validateInput(coord Coordinates, radiusKm float64, edges int) error {
	if edges < 3 {
		return fmt.Errorf("Edges must be greater than 2")
	}

	if radiusKm <= 0 {
		return fmt.Errorf("Radius must be greater than 0")
		// TODO menos del radio de la tierra?
	}

	if coord.Lat < -90 || coord.Lat > 90 {
		return fmt.Errorf("Latitude must be greater than -90 and less than 90")
	}

	if coord.Long < -180 || coord.Long > 180 {
		return fmt.Errorf("Longitude must be greater than -180 and less than 180")
	}
	return nil
}

// Convert circle(radius) to polygon.
// Algorithm source: https://www.movable-type.co.uk/scripts/latlong.html
func polygonFromRadius(coord Coordinates, radiusKm float64, edges int) []Coordinates {

	dR := radiusKm / earthRadiusKM
	res := make([]Coordinates, 0, edges)

	for i := 0; i < edges; i++ {

		brng := float64(2*i) * math.Pi / float64(edges)
		lat1 := coord.Lat * math.Pi / 180
		lon1 := coord.Long * math.Pi / 180

		la := math.Asin(math.Sin(lat1)*math.Cos(dR) +
			math.Cos(lat1)*math.Sin(dR)*math.Cos(brng))

		lo := lon1 + math.Atan2(math.Sin(brng)*math.Sin(dR)*math.Cos(lat1),
			math.Cos(dR)-math.Sin(lat1)*math.Sin(la))

		res = append(res, Coordinates{Long: lo * 180 / math.Pi, Lat: la * 180 / math.Pi})
	}

	return res
}

type Request struct {
	ID          string
	Coordinates Coordinates
	Radius      float64
	Edges       int
}

type Response struct {
	ID      string
	Polygon []Coordinates
}

type Reader interface {
	Read() (Request, error)
}

type Writer interface {
	Write(Response) error
}

func polygonFromReq(req Request) (Response, error) {
	coord, err := FromRadius(req.Coordinates, req.Radius, req.Edges)
	return Response{ID: req.ID, Polygon: coord}, err
}

func FromRadiusIO(in Reader, out Writer) error {
	for {
		req, err := in.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		polResp, err := polygonFromReq(req)
		if err != nil {
			return err
		}

		err = out.Write(polResp)
		if err != nil {
			return err
		}
	}
	return nil
}
