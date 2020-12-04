package main

import (
	"context"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"time"
)

func main() {
	//读取yaml文件
	v := viper.New()
	//设置读取的配置文件
	v.SetConfigName("config")
	//添加读取的配置文件路径
	v.AddConfigPath("./")
	//windows环境下为%GOPATH，linux环境下为$GOPATH
	v.AddConfigPath("$GOPATH/src/learnDemo/viper")
	//设置配置文件类型
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("err:%s\n", err)
	}
	//创建一个信道等待关闭（模拟服务器环境）
	ctx, cancel := context.WithCancel(context.Background())
	//设置监听回调函数
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("config is change :%s \n", e.String())
		cancel()
	})
	//开始监听
	v.WatchConfig()
	//信道不会主动关闭，可以主动调用cancel关闭
	<-ctx.Done()
	fmt.Printf(
		`
		TimeStamp:%s
		CompanyInfomation.Name:%s
		CompanyInfomation.Department:%s `,
		v.Get("TimeStamp"),
		v.Get("CompanyInfomation.Name"),
		v.Get("CompanyInfomation.Department"),
	)

	/*
		result:
		TimeStamp:2018-10-18 10:09:23
		CompanyInfomation.Name:Sunny
		CompanyInfomation.Department:[Finance Design Program Sales]
	*/
	//反序列化
	parseYaml(v)
}

type CompanyInfomation struct {
	Name                 string
	MarketCapitalization int64
	EmployeeNum          int64
	Department           []interface{}
	IsOpen               bool
	Time                 time.Duration
}

type YamlSetting struct {
	TimeStamp         string
	Address           string
	Postcode          int64
	CompanyInfomation CompanyInfomation
}

func parseYaml(v *viper.Viper) {
	var yamlObj YamlSetting
	if err := v.Unmarshal(&yamlObj); err != nil {
		fmt.Printf("err:%s", err)
	}
	dtime := yamlObj.CompanyInfomation.Time
	fmt.Printf("Time [%v],[%d],[%T]\n", dtime, dtime, dtime)
	fmt.Println(yamlObj)
	/*
		result:
		{2018-10-18 10:09:23 Shenzhen 518000 {Sunny 50000000 200 [Finance Design Program Sales] false}}
	*/
}
