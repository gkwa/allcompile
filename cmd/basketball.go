package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	ball "github.com/taylormonacelli/allcompile/basketball"
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
	var ballFactory ball.BallFactory

	ballFactory = ball.NewBasketballFactory()
	b1 := ballFactory.CreateBasketBall()

	fmt.Println()
	fmt.Println("Basketball 1:")
	fmt.Println(b1)

	ballFactory = ball.NewBasketballFactory(
		ball.WithBrand("Wilson"),
		ball.WithColor("Red"),
		ball.WithWeight(20),
	)

	b2 := ballFactory.CreateBasketBall()

	fmt.Println()
	fmt.Println("Basketball 2:")
	fmt.Println(b2)
}
