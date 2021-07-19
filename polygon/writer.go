package polygon

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io"
)

// TODO: GeoJson Writer

type kmlWriter struct {
	encoder *xml.Encoder
	dwriter *bufio.Writer
}

type xmlResp struct {
	XMLName     xml.Name `xml:"Placemark"`
	Name        string   `xml:"name"`
	Coordinates xmlCoord `xml:"Polygon>outerBoundaryIs>LinearRing>coordinates"`
}
type xmlCoord struct {
	XMLName   xml.Name `xml:"coordinates"`
	Formatted string   `xml:",innerxml"`
}

func NewKMLRespWriter(w io.Writer) *kmlWriter {
	e := xml.NewEncoder(w)
	e.Indent("  ", "  ")
	dw := bufio.NewWriter(w)
	return &kmlWriter{encoder: e, dwriter: dw}
}

func (k *kmlWriter) Write(resp Response) error {
	f := k.format(resp.Polygon)
	r := xmlResp{Name: resp.ID, Coordinates: xmlCoord{Formatted: f}}
	// TODO: TEST err
	return k.encoder.Encode(r)
}

const (
	header = `<?xml version="1.0" encoding="UTF-8"?>` + "\n" +
		`<kml xmlns="http://www.opengis.net/kml/2.2">` + "\n" +
		`  <Document>` + "\n"
	footer = `  </Document>` + "\n" +
		`</kml>`
)

func (k *kmlWriter) WriteHeader() error {
	_, e := k.dwriter.WriteString(header)
	e = k.dwriter.Flush()
	// TODO: TEST err
	return e
}

func (k *kmlWriter) WriteFooter() error {
	_, e := k.dwriter.WriteString(footer)
	e = k.dwriter.Flush()
	// TODO: TEST err
	return e
}

func (k *kmlWriter) format(cs []Coordinates) string {
	var r string
	for _, c := range cs {
		r += fmt.Sprintf(" %.8f,%.8f", c.Long, c.Lat)
	}
	r += fmt.Sprintf(" %.8f,%.8f", cs[0].Long, cs[0].Lat)
	return r
}
