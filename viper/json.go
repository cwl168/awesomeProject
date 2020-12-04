package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// 定义配置文件解析后的结构
type MySQLConfig struct {
	URL      string
	Username string
	Password string
}

type Config struct {
	Port  int
	MySQL MySQLConfig
}

func main() {
	var config Config

	// ReadFile 函数会读取文件的全部内容，返回一个字节数组
	data, err := ioutil.ReadFile("/Users/caoweilin/go/src/learnDemo/viper/config.json")
	if err != nil {
		panic(err)
	}

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
	fmt.Println(config)
}
