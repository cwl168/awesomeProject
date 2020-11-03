package cmd

import (
	"log"

	"learnDemo/tour/internal/json2struct"

	"github.com/spf13/cobra"
)

var jsonCmd = &cobra.Command{
	Use:   "json",
	Short: "json转换和处理",
	Long:  "json转换和处理",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var json2structCmd = &cobra.Command{
	Use:   "struct",
	Short: "json转换",
	Long:  "json转换",
	Run: func(cmd *cobra.Command, args []string) {
		parser, err := json2struct.NewParser(str)
		if err != nil {
			log.Fatalf("json2struct.NewParser err: %v", err)
		}
		content := parser.Json2Struct()
		log.Printf("输出结果: %s", content)
	},
}

//go run tour/main.go json struct --str="{\"code\":200,\"data\":{\"lists\":[{\"id\":9,\"created_on\":1595557131,\"modified_on\":1595557131,\"deleted_on\":0,\"name\":\"游戏\",\"created_by\":\"123\",\"modified_by\":\"\",\"state\":1}],\"total\":1},\"msg\":\"ok\"}"
/**
2020/11/03 18:54:07 输出结果: type Tour struct {
    Code float64
    Data map[string]interface {}
    Msg string
}
*/
func init() {
	jsonCmd.AddCommand(json2structCmd)
	json2structCmd.Flags().StringVarP(&str, "str", "s", "", "请输入json字符串")
}
