/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all tasks",
	Long:  `list all existing tasks in Bolt DB`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := bolt.Open("tasks.db", 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		db.View(func(tx *bolt.Tx) error {

			b := tx.Bucket([]byte("Undone"))

			c := b.Cursor()
			i := 1
			for k, v := c.First(); k != nil; k, v = c.Next() {
				fmt.Printf("Task %d %s: %s\n", i, k, v)
				i++
			}

			return nil
		})
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
