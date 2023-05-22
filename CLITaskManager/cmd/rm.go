/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "remove task",
	Long:  `Completely delete task from undone list`,
	Run: func(cmd *cobra.Command, args []string) {

		db, err := bolt.Open("tasks.db", 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		err = db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("Undone"))
			c := b.Cursor()
			i := 1
			var key, value []byte
			num, _ := strconv.Atoi(args[0])

			for k, v := c.First(); k != nil; k, v = c.Next() {
				if num == i {
					key = make([]byte, len(k))
					value = make([]byte, len(v))
					copy(key, k)
					copy(value, v)
					break
				}
				i++
			}
			_ = b.Delete(key)
			fmt.Printf("Deleted: %s\n", string(value))
			return err
		})

	},
}

func init() {
	rootCmd.AddCommand(rmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
