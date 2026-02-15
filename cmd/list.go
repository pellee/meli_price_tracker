package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list [-i identifier]",
	Short: "list the elements to track.",
	Long:  `list the elements to track. Can list all the elements or only show a specific item of the list.`,
	Run:   listCommand,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func listCommand(cmd *cobra.Command, args []string) {
	for _, item := range listOfItemsTracked {
		fmt.Println(item.Id)
		fmt.Println(item.Name)
		for _, link := range item.Links {
			fmt.Println(link.Link)
			fmt.Println(link.Price)
		}
	}

	os.Exit(1)
}
