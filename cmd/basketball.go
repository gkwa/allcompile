package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/taylormonacelli/allcompile/basketball"
)

// basketballCmd represents the basketball command
var basketballCmd = &cobra.Command{
	Use:   "basketball",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		runBasketball()
	},
}

func init() {
	rootCmd.AddCommand(basketballCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// basketballCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// basketballCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runBasketball() {
	var ballFactory basketball.BallFactory

	// Create a basketball with default properties
	ballFactory = basketball.NewBasketballFactory()
	b1 := ballFactory.CreateBasketBall()

	fmt.Println()
	fmt.Println("Basketball 1:")
	fmt.Println(b1)

	// Create a basketball with default properties
	ballFactory = basketball.NewBasketballFactory(
		basketball.WithSize(10),
		basketball.WithColor("Red"),
		basketball.WithWeight(20),
	)

	// Create a basketball with the specified custom properties
	b2 := ballFactory.CreateBasketBall()

	// Use basketballInstance as needed
	fmt.Println()
	fmt.Println("Basketball 2:")
	fmt.Println(b2)
}
