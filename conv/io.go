package conv

import (
	"io"
	"poly-go/polygon"
)

type pipe struct {
	in          pReader
	out         pWriter
	onceWritten bool
}

type pReader interface {
	read() (request, error)
}

type pWriter interface {
	writeHeader() error
	write(response) error
	writeFooter() error
}
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

func (p pipe) CircleToPolygon() error {
	for {
		req, err := p.in.read()
		if err != nil {
			if err == io.EOF {
				//TODO: err
				if p.onceWritten {
					p.out.writeFooter()
				}
				break
			}
			return err
		}

		polResp, err := p.polygonFromReq(req)
		if err != nil {
			return err
		}

		if !p.onceWritten {
			//TODO: err
			p.out.writeHeader()
			p.onceWritten = true

		}

		err = p.out.write(polResp)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p pipe) polygonFromReq(req request) (response, error) {
	coord, err := polygon.FromRadius(req.coordinates, req.radius, req.edges)
	return response{id: req.id, polygon: coord}, err
}
