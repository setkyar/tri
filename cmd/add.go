package cmd

import (
	"fmt"
	"log"

	"github.com/setkyar/tri/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Long:  `Add will create a new todo item to the list`,
	Run:   addRun,
}

var priority int

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority: 1, 2, 3")
}

func addRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("datafile"))

	if err != nil {
		log.Printf("%v", err)
	}

	for _, x := range args {
		item := todo.Item{Text: x}
		item.SetPriority(priority)

		items = append(items, item)
	}

	err = todo.SaveItems(viper.GetString("datafile"), items)
	if err != nil {
		fmt.Printf("%v", err)
	}
}
