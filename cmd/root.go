package cmd

import (
	"fmt"
	"os"
	"poly-go/conv"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "poly-go",
	Short: "Poly-go converts radius and coordinates from CSV to polygon in format KML or GeoJson",
	Long:  `Poly-go converts radius and coordinates from CSV to polygon in format to KML or GeoJson. For more information...`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: FIRST try
		fmt.Println("This is a test")
		c := conv.NewCsvToKmlConverter(strings.NewReader("ID1, 1.1, 2.2, 1, 4\n ID2, 1.1, 2.2, 111, 50"), os.Stdout)
		err := c.CircleToPolygon()
		if err != nil {
			fmt.Println(err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
