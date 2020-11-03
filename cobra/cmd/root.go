package cmd

import (
	"github.com/spf13/cobra"
	"learnDemo/cobra/imp"
)

var name string
var age int
var rootCmd = &cobra.Command{
	Use:   "demo",
	Short: "this is a demo",
	Long:  `this is a demo long `,
	Run: func(cmd *cobra.Command, args []string) {
		if len(name) == 0 {
			cmd.Help()
			return
		}
		imp.Show(name, age)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().StringVarP(&name, "name", "n", "", "person's name")
	rootCmd.Flags().IntVarP(&age, "age", "a", 0, "Person's age")
}

//go run cobra/main.go  -n=cwl -a=10
