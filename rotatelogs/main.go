package main

import (
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	writer, _ := rotatelogs.New(
		"log"+".%Y%m%d%H%M",
		rotatelogs.WithLinkName("log"),              // 生成软链，指向最新日志文件
		rotatelogs.WithRotationTime(time.Second*60), // 日志切割时间间隔
	)

	logrus.SetOutput(writer)

	f := func() {
		for i := 0; i < 100; i++ {
			logrus.WithFields(logrus.Fields{
				"animal": "walrus",
				"number": i,
			}).Info("A walrus appears")
			time.Sleep(time.Second)
		}
	}

	for i := 0; i < 10; i++ {
		go f()
	}
	f()
	time.Sleep(time.Second)
}
