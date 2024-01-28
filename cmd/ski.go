package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/taylormonacelli/allcompile/ski"
)

// skiCmd represents the ski command
var skiCmd = &cobra.Command{
	Use:   "ski",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		runSki()
	},
}

func init() {
	rootCmd.AddCommand(skiCmd)
}

func runSki() {
	prototype := ski.NewSkiPrototype()

	ski1 := prototype.Clone()
	ski1.Brand = "Atomic"
	ski1.Color = "Red"
	ski1.Weight = 180

	fmt.Println()
	fmt.Println("Ski 1:")
	fmt.Println(ski1)

	ski1.SkiDownhill()
	ski1.Stop()

	ski2 := prototype.Clone()
	ski2.Color = "Blue"

	fmt.Println()
	fmt.Println("Ski 2:")
	fmt.Println(ski2)

	ski2.SkiDownhill()
	ski2.Stop()
}
