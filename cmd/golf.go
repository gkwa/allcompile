package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/taylormonacelli/allcompile/golf"
)

var golfCmd = &cobra.Command{
	Use:   "golf",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		runGolf()
	},
}

func init() {
	rootCmd.AddCommand(golfCmd)
}

func runGolf() {
	printClub("Golf Club 1:", golf.NewGolfClubBuilder().Brand("Titleist").Build())
	printClub("Golf Club 2:", golf.NewGolfClubBuilder().Build())
	printClub("Golf Club 3:", golf.NewGolfClubBuilder().Type("Putter").Length(35).Build())

	// Create a list of golf clubs
	clubs := []*golf.GolfClub{
		golf.NewGolfClubBuilder().Brand("Titleist").Build(),
		golf.NewGolfClubBuilder().Build(),
		golf.NewGolfClubBuilder().Type("Putter").Length(35).Build(),
	}

	// Print each golf club
	for i, club := range clubs {
		printClub(fmt.Sprintf("Golf Club %d:", i+1), club)
	}
}

func printClub(label string, club *golf.GolfClub) {
	fmt.Println()
	fmt.Println(label)
	fmt.Println(club)
}
