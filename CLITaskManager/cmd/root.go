/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/boltdb/bolt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "CLITaskManager",
	Short: "Simple TODO list`",
	Long:  `TODO list with add, list and do functions that write your tasks in Bolt DB`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//Run: func(cmd *cobra.Command, args []string) {
	//
	//},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	db, err := bolt.Open("tasks.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Undone"))
		if b == nil {
			_, err = tx.CreateBucket([]byte("Undone"))
		}
		b = tx.Bucket([]byte("Done"))
		if b == nil {
			_, err = tx.CreateBucket([]byte("Done"))
		}
		return err
	})
	db.Close()
	err = rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.CLITaskManager.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}
