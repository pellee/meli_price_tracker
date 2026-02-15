package cmd

import (
	"github.com/spf13/cobra"
)

var (
	linkToTrack string
	newCmd      = &cobra.Command{
		Use:   "new [item name] [-l --link link to start tracking]",
		Short: "crete a new item.",
		Long:  `create a new item in the list of prices to track.`,
		Run:   newCommand,
	}
)

func init() {
	newCmd.PersistentFlags().StringVarP(&linkToTrack, "link", "l", "", "link to track.")
	rootCmd.AddCommand(newCmd)
}

func newCommand(cmd *cobra.Command, args []string) {
	item := itemTracked{
		Id:   listOfItemsTracked[len(listOfItemsTracked)-1].Id + 1,
		Name: args[0],
	}

	if linkToTrack != "" {
		item.Links = append(item.Links, link{Link: linkToTrack, Price: 0})
	}

	listOfItemsTracked = append(listOfItemsTracked, item)
}
