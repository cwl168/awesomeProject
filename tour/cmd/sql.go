package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"learnDemo/tour/internal/sql2struct"
)

var username string
var password string
var host string
var charset string
var dbType string
var dbName string
var tableName string

//go run tour/main.go sql  struct --username=root --password=root --db=blog --table=blog_tag
var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "sql转换和处理",
	Long:  "sql转换和处理",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var sql2structCmd = &cobra.Command{
	Use:   "struct",
	Short: "sql转换",
	Long:  "sql转换",
	Run: func(cmd *cobra.Command, args []string) {
		dbInfo := &sql2struct.DBInfo{
			DBType:   dbType,
			Host:     host,
			UserName: username,
			Password: password,
			Charset:  charset,
		}
		dbModel := sql2struct.NewDBModel(dbInfo)
		//连接数据库
		err := dbModel.Connect()
		if err != nil {
			log.Fatalf("dbModel.Connect err: %v", err)
		}
		//获取表的列信息
		columns, err := dbModel.GetColumns(dbName, tableName)
		if err != nil {
			log.Fatalf("dbModel.GetColumns err: %v", err)
		}

		template := sql2struct.NewStructTemplate()
		//模板中组装表列信息
		templateColumns := template.AssemblyColumns(columns)
		err = template.Generate(tableName, templateColumns)
		if err != nil {
			log.Fatalf("template.Generate err: %v", err)
		}
	},
}

func init() {
	sqlCmd.AddCommand(sql2structCmd)
	sql2structCmd.Flags().StringVarP(&username, "username", "", "", "请输入数据库的账号")
	sql2structCmd.Flags().StringVarP(&password, "password", "", "", "请输入数据库的密码")
	sql2structCmd.Flags().StringVarP(&host, "host", "", "127.0.0.1:3306", "请输入数据库的HOST")
	sql2structCmd.Flags().StringVarP(&charset, "charset", "", "utf8mb4", "请输入数据库的编码")
	sql2structCmd.Flags().StringVarP(&dbType, "type", "", "mysql", "请输入数据库实例类型")
	sql2structCmd.Flags().StringVarP(&dbName, "db", "", "", "请输入数据库名称")
	sql2structCmd.Flags().StringVarP(&tableName, "table", "", "", "请输入表名称")
}

/**
type BlogTag struct {
         // id
         Id     int32   `json:"id"`
         // 标签名称
         Name   string  `json:"name"`
         // 创建时间
         CreatedOn      int32   `json:"created_on"`
         // 创建人
         CreatedBy      string  `json:"created_by"`
         // 修改时间
         ModifiedOn     int32   `json:"modified_on"`
         // 修改人
         ModifiedBy     string  `json:"modified_by"`
         // deleted_on
         DeletedOn      int32   `json:"deleted_on"`
         // 状态 0为禁用、1为启用
         State  int8    `json:"state"`
}
func (model BlogTag) TableName() string {
        return "blog_tag"
}
*/
