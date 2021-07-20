package conv

import (
	"io"
	"poly-go/polygon"
)

type request struct {
	id          string
	coordinates polygon.Coordinates
	radius      float64
	edges       int
}

type response struct {
	id      string
	polygon []polygon.Coordinates
}

type reader interface {
	read() (request, error)
}

type writer interface {
	write(response) error
}

func polygonFromReq(req request) (response, error) {
	coord, err := polygon.FromRadius(req.coordinates, req.radius, req.edges)
	return response{id: req.id, polygon: coord}, err
}

func FromRadiusIO(in reader, out writer) error {
	for {
		req, err := in.read()
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

		err = out.write(polResp)
		if err != nil {
			return err
		}
	}
	return nil
}
