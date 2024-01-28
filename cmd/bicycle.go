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
		run()
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

func run() {
	bike1 := bicycle.NewBicycle(
		bicycle.WithBrand("Giant"),
		bicycle.WithColor("Blue"),
		bicycle.WithWeight(10),
	)

	fmt.Println()
	fmt.Println("Bicycle 1:")
	fmt.Println(bike1)

	bike1.Pedal()
	bike1.Brake()

	bike2 := bicycle.NewBicycle(
		bicycle.WithColor("Green"),
	)

	fmt.Println()
	fmt.Println("Bicycle 2:")
	fmt.Println(bike2)

	bike2.Pedal()
	bike2.Brake()

	bike3 := &bicycle.Bicycle{
		Brand:  "Specialized",
		Color:  "Red",
		Weight: 12,
	}

	fmt.Println()
	fmt.Println("Bicycle 3:")
	bike3.Brake()
	fmt.Println(bike3)

	bike4 := &bicycle.Bicycle{
		Brand: "SpeedX",
	}

	fmt.Println()
	fmt.Println("Bicycle 4:")
	bike4.Pedal()
	fmt.Println(bike4)

	bike5 := bicycle.NewBicycle(
		bicycle.WithBrand("Giant"),
		bicycle.WithWeight(10),
	)

	fmt.Println()
	fmt.Println("Bicycle 5:")
	bike5.Pedal()
	fmt.Println(bike5)

	bike6 := bicycle.NewBicycle()

	fmt.Println()
	fmt.Println("Bicycle 6:")
	bike6.Pedal()
	fmt.Println(bike5)
}
