/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/setkyar/tri/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	doneOpt bool
	allOpt  bool
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the todos",
	Long:  `Listing todo list.`,
	Run:   listRun,
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVar(&doneOpt, "done", false, "Show 'Done' Todos")
	listCmd.Flags().BoolVar(&allOpt, "all", false, "Show all Todods")
}

func listRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("datafile"))

	if err != nil {
		log.Printf("%v", err)
	}

	sort.Sort(todo.ByPri(items))

	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
	for _, i := range items {

		if allOpt || i.Done == doneOpt {
			fmt.Fprintln(w, i.Label()+"\t"+i.PrettyDone()+"\t"+i.PrettyP()+"\t"+i.Text+"\t")
		}
	}

	w.Flush()
}
