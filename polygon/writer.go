package polygon

import (
	"encoding/xml"
	"fmt"
	"io"
)

// TODO: GeoJson Writer

type kmlWriter struct {
	encoder *xml.Encoder
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

func NewKMLRespWriter(w io.Writer) Writer {
	enc := xml.NewEncoder(w)
	enc.Indent("", "    ")
	return &kmlWriter{encoder: enc}
}

func (k *kmlWriter) Write(resp Response) error {
	f := k.format(resp.Polygon)
	r := xmlResp{Name: resp.ID, Coordinates: xmlCoord{Formatted: f}}
	// TODO: TEST err
	return k.encoder.Encode(r)
}

func (k *kmlWriter) format(cs []Coordinates) string {
	var r string
	for _, c := range cs {
		r += fmt.Sprintf(" %v,%v", c.Long, c.Lat)
	}
	r += fmt.Sprintf(" %v,%v ", cs[0].Long, cs[0].Lat)
	return r
}
