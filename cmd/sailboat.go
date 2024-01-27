/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/taylormonacelli/allcompile/sailboat"
)

// sailboatCmd represents the sailboat command
var sailboatCmd = &cobra.Command{
	Use:   "sailboat",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sailboat called")
		runSailboat()
	},
}

func init() {
	rootCmd.AddCommand(sailboatCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sailboatCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sailboatCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runSailboat() {
	// Using the SailboatBuilder to create instances with specified properties
	sailboat1 := sailboat.NewSailboatBuilder().
		Brand("Giant").
		Color("Blue").
		Build()

	sailboat2 := sailboat.NewSailboatBuilder().
		Color("Green").
		Build()

	// Printing information about the Sailboats
	fmt.Println("Sailboat 1:")
	fmt.Println(sailboat1.String())

	fmt.Println()

	fmt.Println("Sailboat 2:")
	fmt.Println(sailboat2.String())

	// Using the SailboatBuilder to create instances with specified properties
	boat1 := sailboat.NewSailboatBuilder().
		Brand("Sunfish").
		Color("White").
		Length(14).
		Build()

	boat2 := sailboat.NewSailboatBuilder().
		Brand("Hobie Cat").
		Length(16).
		Build()

	// Printing information about the sailboats
	fmt.Println("\nSailboat 1:")
	fmt.Println(boat1.String())

	fmt.Println()

	fmt.Println("Sailboat 2:")
	fmt.Println(boat2.String())
}
