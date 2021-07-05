package polygon

import "fmt"

// TODO: KML and GeoJson Writer . Just for test pourpuses
type ConsoleWriter struct {
}

func (c *ConsoleWriter) Write(resp Response) error {
	fmt.Println(resp)
	return nil
}
