package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/taylormonacelli/allcompile/kite"
)

// kiteCmd represents the kite command
var kiteCmd = &cobra.Command{
	Use:   "kite",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		k := kite.NewKite(
			kite.WithBrand("KiteCo"),
			kite.WithColor("Red"),
			kite.WithLength(30),
		)

		fmt.Println()
		fmt.Println("Kite 1:")
		fmt.Println(k)
	},
}

func init() {
	rootCmd.AddCommand(kiteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kiteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kiteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
