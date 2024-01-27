/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/taylormonacelli/allcompile/bicycle"
)

// bicycleCmd represents the bicycle command
var bicycleCmd = &cobra.Command{
	Use:   "bicycle",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Creating a new Bicycle instance with specified properties
		bike1 := &bicycle.Bicycle{
			Brand:  "Giant",
			Color:  "Blue",
			Weight: 10,
		}

		// Creating a new generic Bicycle instance with default values
		bike2 := bicycle.NewBicycle()

		// Displaying information about the first bicycle
		fmt.Println("Bicycle 1:")
		fmt.Println(bike1.String())

		// Performing some actions on the first bicycle
		bike1.Pedal()
		bike1.Brake()

		fmt.Println()

		// Displaying information about the second bicycle
		fmt.Println("Bicycle 2:")
		fmt.Println(bike2.String())

		// Performing some actions on the second bicycle
		bike2.Pedal()
		bike2.Brake()

	},
}

func init() {
	rootCmd.AddCommand(bicycleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// bicycleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// bicycleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
