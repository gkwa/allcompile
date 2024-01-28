package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/taylormonacelli/allcompile/hockey"
)

// hockeyCmd represents the hockey command
var hockeyCmd = &cobra.Command{
	Use:   "hockey",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		runHockey()
	},
}

func init() {
	rootCmd.AddCommand(hockeyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hockeyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hockeyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runHockey() {
	stick1 := hockey.NewHockeyStick(
		hockey.HockeyStick{
			Brand:  "Bauer",
			Length: 63,
			Color:  "Blue",
		},
	)

	fmt.Println()
	fmt.Println("HockeyStick 1:")
	fmt.Println(stick1)

	stick1.Shoot()
	stick1.Pass()

	stick2 := hockey.NewHockeyStick(
		hockey.HockeyStick{
			Brand: "CCM",
		},
	)

	fmt.Println()
	fmt.Println("HockeyStick 2:")
	fmt.Println(stick2)

	stick2.Shoot()
	stick2.Pass()

	stick3 := &hockey.HockeyStick{}

	fmt.Println()
	fmt.Println("HockeyStick 3:")
	stick3.Pass()
	fmt.Println(stick3)
}
