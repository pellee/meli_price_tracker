package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

const DEFAULT_PATH string = "/home/pelle/mpt_trackings.json"

var listOfItemsTracked []itemTracked

type itemTracked struct {
	Id    int32  `json:"id"`
	Name  string `json:"name"`
	Links []link `json:"links"`
}

type link struct {
	Link  string  `json:"link"`
	Price float32 `json:"price"`
}

var rootCmd = &cobra.Command{
	Use:   "mpt",
	Short: "MPT (Meli Price Tracker) is simple way to track your futures buys",
	Long: `A simple way of tracking future buys.
	just keep it simple idiot.`,
	PersistentPreRun:  loadList,
	PersistentPostRun: saveList,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func showError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func loadList(cmd *cobra.Command, args []string) {
	json_file, err := os.ReadFile(DEFAULT_PATH)
	showError(err)

	err = json.Unmarshal(json_file, &listOfItemsTracked)
	showError(err)
}

func saveList(cmd *cobra.Command, args []string) {
	data, err := json.MarshalIndent(listOfItemsTracked, "", " ")
	showError(err)

	os.WriteFile(DEFAULT_PATH, data, 0644)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
