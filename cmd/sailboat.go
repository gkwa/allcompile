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
	boat1 := sailboat.NewSailboat()
	fmt.Println("Sailboat 1:")
	fmt.Println(boat1.String())

	boat2 := sailboat.NewSailboat(
		sailboat.WithLength(100),
	)
	fmt.Println()
	fmt.Println("Sailboat 2:")
	fmt.Println(boat2.String())

	boat3 := sailboat.NewSailboatBuilder().
		Brand("Hobie Cat").
		Build()
	fmt.Println()
	fmt.Println("Sailboat 3:")
	fmt.Println(boat3.String())
}
