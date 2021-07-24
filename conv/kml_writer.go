package conv

import (
	"encoding/xml"
	"fmt"
	"io"
	"poly-go/polygon"
)

type kmlWriter struct {
	encoder *xml.Encoder
	wr      io.Writer
}

const (
	header = `<?xml version="1.0" encoding="UTF-8"?>` + "\n" +
		`<kml xmlns="http://www.opengis.net/kml/2.2">` + "\n" +
		"  <Document>\n"
	footer = "\n" +
		"  </Document>\n" +
		"</kml>"
)

type xmlResp struct {
	XMLName     xml.Name `xml:"Placemark"`
	Name        string   `xml:"name"`
	Coordinates xmlCoord `xml:"Polygon>outerBoundaryIs>LinearRing>coordinates"`
}
type xmlCoord struct {
	XMLName   xml.Name `xml:"coordinates"`
	Formatted string   `xml:",innerxml"`
}

func newKMLRespWriter(w io.Writer) *kmlWriter {
	e := xml.NewEncoder(w)
	e.Indent("    ", "  ")
	return &kmlWriter{encoder: e, wr: w}
}

func (k *kmlWriter) write(resp response) error {
	f := k.format(resp.polygon)
	xr := xmlResp{Name: resp.id, Coordinates: xmlCoord{Formatted: f}}
	// TODO: TEST err
	k.encoder.Encode(xr)
	return nil
}

func (k *kmlWriter) writeHeader() error {
	// TODO: TEST err
	k.wr.Write([]byte(header))
	return nil
}

func (k *kmlWriter) writeFooter() error {
	// TODO: TEST err
	k.wr.Write([]byte(footer))
	return nil
}

func (k *kmlWriter) format(cs []polygon.Coordinates) string {
	var r string
	for _, c := range cs {
		r += fmt.Sprintf(" %.8f,%.8f", c.Long, c.Lat)
	}
	r += fmt.Sprintf(" %.8f,%.8f", cs[0].Long, cs[0].Lat)
	return r
}
