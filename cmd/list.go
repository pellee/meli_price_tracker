package cmd

import (
	"os"
	"text/tabwriter"

	"meli_price_tracker/utils"

	"github.com/spf13/cobra"
)

var (
	id      int32
	allFlag bool
)

var listCmd = &cobra.Command{
	Use:   "list [id optionally] [-a --all]",
	Short: "list the elements to track.",
	Long:  `list the elements to track. Can list all the elements or only show a specific item of the list.`,
	Run:   listCommand,
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVarP(&allFlag, "all", "a", false, "to get all the information")
	listCmd.Flags().Int32VarP(&id, "id", "i", 0, "list a specific item in the list.")
}

func listAll() {
	writter := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	utils.AddItem(writter, []any{"Id", "Name", "Link", "Price"})

	for _, item := range listOfItemsTracked {
		if id == -1 {
			if len(item.Links) > 0 {
				for _, linkItem := range item.Links {
					utils.AddItem(writter, []any{item.Id, item.Name, linkItem.Link, linkItem.Price})
				}
			} else {
				utils.AddItem(writter, []any{item.Id, item.Name, "null", ""})
			}
		} else {
			if id == item.Id {
				if len(item.Links) > 0 {
					for _, linkItem := range item.Links {
						utils.AddItem(writter, []any{item.Id, item.Name, linkItem.Link, linkItem.Price})
					}
				} else {
					utils.AddItem(writter, []any{item.Id, item.Name, "null", ""})
				}
			}
		}
	}
	writter.Flush()
}

func listAllLowerPrice() {
	writter := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	utils.AddItem(writter, []any{"Id", "Name", "Link", "Price"})

	for _, item := range listOfItemsTracked {
		if id == -1 {
			if len(item.Links) > 0 {
				finalPrice := item.Links[0].Price
				idxFinalPrice := 0
				for i := 1; i < len(item.Links); i++ {
					if item.Links[i].Price < finalPrice {
						finalPrice = item.Links[i].Price
						idxFinalPrice = i
					}
				}
				utils.AddItem(writter, []any{item.Id, item.Name, item.Links[idxFinalPrice].Link, item.Links[idxFinalPrice].Price})
			}
		} else {
			if id == item.Id {
				if len(item.Links) > 0 {
					finalPrice := item.Links[0].Price
					idxFinalPrice := 0
					for i := 1; i < len(item.Links); i++ {
						if item.Links[i].Price < finalPrice {
							finalPrice = item.Links[i].Price
							idxFinalPrice = i
						}
					}
					utils.AddItem(writter, []any{item.Id, item.Name, item.Links[idxFinalPrice].Link, item.Links[idxFinalPrice].Price})
				}
			}
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
