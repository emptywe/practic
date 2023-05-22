/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark task as done",
	Long:  `Mark chosen task as done in Bolt DB`,
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
			fmt.Printf("Done %s %s\n", string(key), string(value))
			b = tx.Bucket([]byte("Done"))
			b.Bucket([]byte("Done"))
			err = b.Put(key, value)
			return err
		})

	},
}

func init() {
	rootCmd.AddCommand(doCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
