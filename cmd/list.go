package cmd

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

// id      int32
var allFlag bool

var listCmd = &cobra.Command{
	Use:   "list [id optionally] [-a --all]",
	Short: "list the elements to track.",
	Long:  `list the elements to track. Can list all the elements or only show a specific item of the list.`,
	Run:   listCommand,
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&allFlag, "all", "a", false, "to get all the information")
}

func formatString(params []string) string {
	var buff bytes.Buffer

	for i := range params {
		buff.WriteString("%v")
		if i+1 != len(params) {
			buff.WriteString("\t")
		}
	}

	buff.WriteString("\n")

	return buff.String()
}

func listAll() {
	writter := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	formatedString := formatString([]string{"Id", "Name", "Link", "Price"})
	fmt.Fprintf(writter, formatedString, "Id", "Name", "Link", "Price")

	for _, item := range listOfItemsTracked {
		if len(item.Links) > 0 {
			for _, linkItem := range item.Links {
				formatedString = formatString([]string{string(item.Id), item.Name, linkItem.Link, strconv.FormatFloat(float64(linkItem.Price), 'f', -1, 32)})
				fmt.Fprintf(writter, formatedString, item.Id, item.Name, linkItem.Link, linkItem.Price)
			}
		} else {
			formatedString = formatString([]string{string(item.Id), item.Name, "null", ""})
			fmt.Fprintf(writter, formatedString, item.Id, item.Name, "null", 0)
		}
	}
	writter.Flush()
}

func listAllLowerPrice() {
	writter := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	formatedString := formatString([]string{"Id", "Name", "Link", "Price"})
	fmt.Fprintf(writter, formatedString, "Id", "Name", "Link", "Price")

	for _, item := range listOfItemsTracked {
		if len(item.Links) > 0 {
			finalPrice := item.Links[0].Price
			idxFinalPrice := 0
			for i := 1; i < len(item.Links); i++ {
				if item.Links[i].Price < finalPrice {
					finalPrice = item.Links[i].Price
					idxFinalPrice = i
				}
			}
			formatedString = formatString([]string{string(item.Id), item.Name, item.Links[idxFinalPrice].Link, strconv.FormatFloat(float64(item.Links[idxFinalPrice].Price), 'f', -1, 32)})
			fmt.Fprintf(writter, formatedString, item.Id, item.Name, item.Links[idxFinalPrice].Link, item.Links[idxFinalPrice].Price)
		}
	}
	writter.Flush()
}

func listCommand(cmd *cobra.Command, args []string) {
	if allFlag {
		listAll()
		os.Exit(1)
	}

	listAllLowerPrice()
	os.Exit(1)
}
