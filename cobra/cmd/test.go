package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

//添加子命令test，在cmd目录下生成test.go
// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "id test",
	Long:  `this is a id test`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test called")
		//go run cobra/main.go  test  --name=cwl
		//go run cobra/main.go  test  -n=cwl
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			fmt.Println("请输入正确的名称")
			return
		}
		fmt.Println("名称：", name, cmd.Name())
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
	testCmd.Flags().StringVarP(&name, "name", "n", "", "person's name")

}
