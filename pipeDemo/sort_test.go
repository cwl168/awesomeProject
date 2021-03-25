package pipeDemo

import (
	"github.com/fortytw2/leaktest"
	"testing"
)

func TestMain0(t *testing.T) {
	main0()
}

//归并排序网络版
func TestMain1(t *testing.T) {
	defer leaktest.Check(t)()
	main2()
}

//文件归并排序
func TestMain3(t *testing.T) {
	defer leaktest.Check(t)()
	main3()
}

//生成一批随机数并写入文件
func TestMain4(t *testing.T) {
	main4()
}
