/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, _ []string) {
		english, _ := cmd.Flags().GetString("english")
		german, _ := cmd.Flags().GetString("german")

		connection, _ := ConnectToDatabase()
		card := Card{English: english, German: german}

		connection.Create(&card)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("english", "e", "", "The english version of the word")
	createCmd.MarkFlagRequired("english")

	createCmd.Flags().StringP("german", "g", "", "The german version of the word")
	createCmd.MarkFlagRequired("german")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
