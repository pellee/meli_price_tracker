package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

type itemTracked struct {
	Id    int32  `json:"id"`
	Name  string `json:"name"`
	Links []link `json:"links"`
}

type link struct {
	Link  string  `json:"link"`
	Price float32 `json:"price"`
}

var listCmd = &cobra.Command{
	Use:   "list [-i identifier]",
	Short: "list the elements to track.",
	Long:  `list the elements to track. Can list all the elements or only show a specific item of the list.`,
	Run:   listCommand,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func showError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func listCommand(cmd *cobra.Command, args []string) {
	json_file, err := os.ReadFile("./test_files/example.json")
	showError(err)

	var payload []itemTracked
	err = json.Unmarshal(json_file, &payload)
	showError(err)
	// result, err := json.MarshalIndent(payload, "", "\t")
	///showError(err)

	fmt.Println(payload[0].Id)
	fmt.Println(payload[0].Name)
	fmt.Println(payload[0].Links)
}
