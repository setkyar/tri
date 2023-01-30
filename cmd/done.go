package cmd

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/setkyar/tri/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:     "done",
	Aliases: []string{"do"},
	Short:   "Mark Item as Done",
	Long:    `Mark item as done`,
	Run:     doneRun,
}

func init() {
	rootCmd.AddCommand(doneCmd)
}

func doneRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("datafile"))

	if err != nil {
		log.Fatalf("Read items : %v\n", err)
	}

	i, err := strconv.Atoi(args[0])

	if err != nil {
		log.Fatalln(args[0], "is not a valid label\n", err)
	}

	if i > 0 && i < len(items) {
		items[i-1].Done = true

		fmt.Printf("%q %v\n", items[i-1].Text, "marked done")

		sort.Sort(todo.ByPri(items))
		todo.SaveItems(viper.GetString("datafile"), items)
	} else {
		log.Println(i, "doesn't match any items")
	}
}
