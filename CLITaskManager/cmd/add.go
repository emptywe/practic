/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adding task",
	Long:  `Adding task that you need into Bolt DB`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := bolt.Open("tasks.db", 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		err = db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("Undone"))
			// Generate ID for the user.
			// This returns an error only if the Tx is closed or not writeable.
			// That can't happen in an Update() call so I ignore the error check.
			key := time.Now().Format("2006-01-02 15-04-05")
			if err != nil {
				return err
			}

			return b.Put([]byte(key), []byte(strings.Join(args, " ")))
		})
		fmt.Printf("Added: %s \n", strings.Join(args, " "))
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
