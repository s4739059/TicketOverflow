package cmd

import (
	"encoding/json"
	"fmt"
	"hamilton/service"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

// seatingCmd represents the seating command
var seatingCmd = &cobra.Command{
	Use:   "seating",
	Short: "Generate a seating plan SVG",
	Long: `Generate a seating plan SVG with the given input file as describe by:

	{
		"id": "12345678-1234-1234-1234-123456789012",
		"name": "Example Concert",
		"date": "2021-01-01",
		"venue": "Example Venue",
		"seats": {
			"max": 100,
			"purchased": 56
		}
	}
`,
	Run: func(cmd *cobra.Command, args []string) {
		// get the input and output command line arguments
		input, _ := cmd.Flags().GetString("input")
		output, _ := cmd.Flags().GetString("output")

		rawInfo, _ := ioutil.ReadFile(input)
		var info service.Concert
		err := json.Unmarshal(rawInfo, &info)
		if err != nil {
			errorAndClose(err, output)
		}

		pencil := service.NewDrawer()
		concert, err := pencil.DrawConcert(info)
		if err != nil {
			errorAndClose(err, output)
		}

		f, err := os.Create(fmt.Sprintf("%s.svg", output))
		if err != nil {
			errorAndClose(err, output)
		}
		defer f.Close()

		_, err = f.WriteString(concert)
		if err != nil {
			errorAndClose(err, output)
		}
	},
}

func init() {
	generateCmd.AddCommand(seatingCmd)

	seatingCmd.Flags().StringP("input", "i", "input.json", "Path of the input file")
	seatingCmd.Flags().StringP("output", "o", "output", "Path of the output file without extension")
}
