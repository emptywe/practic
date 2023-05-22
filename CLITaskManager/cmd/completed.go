/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"time"

	"github.com/spf13/cobra"
)

// complitedCmd represents the complited command
var complitedCmd = &cobra.Command{
	Use:   "completed",
	Short: "list today's completed tasks",
	Long:  `Show all tasks that were completed today`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := bolt.Open("tasks.db", 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		db.View(func(tx *bolt.Tx) error {
			// Assume bucket exists and has keys
			b := tx.Bucket([]byte("Done"))

			c := b.Cursor()
			i := 1
			for k, v := c.First(); k != nil; k, v = c.Next() {
				t, _ := time.Parse("2006-01-02 15-04-05", string(k))
				if t.Day() == time.Now().Day() && t.Month() == time.Now().Month() && t.Year() == time.Now().Year() {
					fmt.Printf("Task %d %s: %s\n", i, k, v)
					i++
				} else {
					_ = b.Delete(k)
				}
			}

			return nil
		})
	},
}

func init() {
	rootCmd.AddCommand(complitedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// complitedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// complitedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
