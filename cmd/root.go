package cmd

import (
	"bufio"
	"log"
	"os"
	"regexp"

	"github.com/emmanuelgautier/domain-scout/scout"
	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/renderer"
	"github.com/olekukonko/tablewriter/tw"
	"github.com/spf13/cobra"
)

func extractDomains(input string) []string {
	regex := regexp.MustCompile(`(?m)(?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?\.)+[a-z0-9][a-z0-9-]{0,61}[a-z0-9]`)
	matches := regex.FindAllString(input, -1)
	return matches
}

var subdomainAvailableCmd = &cobra.Command{
	Use:   "subdomain-available",
	Short: "Check if subdomains are available",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewScanner(os.Stdin)
		var input string
		for reader.Scan() {
			input += reader.Text() + "\n"
		}

		if reader.Err() != nil {
			log.Fatal(reader.Err())
			return
		}

		domains := extractDomains(input)
		availabilities, err := scout.CheckAvailability(cmd.Context(), domains)
		if err != nil {
			log.Fatal(err)
			return
		}

		var data [][]string
		for _, availability := range availabilities {
			if len(availability.Records.Records) == 0 {
				data = append(data, []string{
					availability.Domain,
					"Yes",
					"",
					"",
				})
				continue
			}

			var httpResponse = ""
			if availability.IsRootHTTPReachable != nil {
				httpResponse = availability.IsRootHTTPReachable.String()
			}

			for _, record := range availability.Records.Records {
				data = append(data, []string{
					availability.Domain,
					"No",
					"(" + record.Type + ") " + record.Value,
					httpResponse,
				})
			}
		}

		table := tablewriter.NewTable(os.Stdout,
			tablewriter.WithRenderer(renderer.NewBlueprint(tw.Rendition{
				Settings: tw.Settings{Separators: tw.Separators{BetweenRows: tw.On}},
			})),
			tablewriter.WithConfig(tablewriter.Config{
				Header: tw.CellConfig{Alignment: tw.CellAlignment{Global: tw.AlignCenter}},
				Row: tw.CellConfig{
					Merging:   tw.CellMerging{Mode: tw.MergeBoth},
					Alignment: tw.CellAlignment{Global: tw.AlignLeft},
				},
			}),
		)
		table.Header([]string{"Domain", "Available", "Records", "HTTP Response"})
		err = table.Bulk(data)
		if err != nil {
			log.Fatal(err)
			return
		}

		err = table.Render()
		if err != nil {
			log.Fatal(err)
			return
		}
	},
}

func NewRootCmd() (cmd *cobra.Command) {
	var rootCmd = &cobra.Command{
		Use:   "domain-scout",
		Short: "Scan domains and subdomains",
	}

	rootCmd.AddCommand(subdomainAvailableCmd)

	return rootCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	c := NewRootCmd()

	if err := c.Execute(); err != nil {
		os.Exit(1)
	}
}
