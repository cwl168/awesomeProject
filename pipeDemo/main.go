package main

import (
	"bufio"
	"fmt"
	"learnDemo/pipeline"
	"os"
	"strconv"
)

//生成大的测试文件
const FILE_NAME = "large.in"
const N = 4000

func main00() {

	file, e := os.Create(FILE_NAME)
	if e != nil {
		panic(e)
	}
	defer file.Close()

	p := pipeline.RandomSource(N)
	writer := bufio.NewWriter(file)
	pipeline.WriteSink(writer, p)
	writer.Flush()

	file, e = os.Open(FILE_NAME)
	if e != nil {
		panic(e)
	}
	defer file.Close()

	//p = pipeline.ReaderSource(file, -1)
	p = pipeline.ReaderSource(file, 50)

	//count := 0
	/*for v := range p {
		fmt.Println(v)
		count++
		if count < 100 {
			fmt.Println(v)
		}
	}*/

	//p := pipeline.ArraySource(3, 2, 6, 7, 4)p
	ch := pipeline.InMemSort(p)
	for v := range ch {
		fmt.Println(v)
	}

	//for {
	//	// 因为pipeline 会关闭通道，所以这边收到的数据是0,换成以下写法
	//	//num := <-p
	//	//fmt.Println(num)
	//
	//	// ok 如果channel还在，ok为true，如果channel关闭，ok为false
	//	if num,ok :=<-p;ok{
	//		fmt.Println(num)
	//	}else {
	//		break
	//	}
	//}
	//不断从p收数据，如果没数据，可能会等待

}

//手动指定数组，合并排序
func main0() {
	p := pipeline.Merge(
		pipeline.InMemSort(pipeline.ArraySource(8, 2, 3, 0, 1)),
		pipeline.InMemSort(pipeline.ArraySource(9, 7)),
	)
	for v := range p {
		fmt.Println(v)
	}
}
func main4() {
	//创建文件 生成随机数字 随机数字写入文件
	file, e := os.Create(FILE_NAME)
	if e != nil {
		panic(e)
	}
	defer file.Close()

	p := pipeline.RandomSource(N)
	writer := bufio.NewWriter(file)
	pipeline.WriteSink(writer, p)
	writer.Flush()

	file, e = os.Open(FILE_NAME)
	if e != nil {
		panic(e)
	}
	defer file.Close()

	fmt.Println("读文件结果如下：")
	p1 := pipeline.ReaderSource(file, -1)
	for v1 := range p1 {
		fmt.Println(v1)
	}
}
func main3() {
	p := createPipeline("large.in", 320, 4) //4相当于4个结点，跟CPU核数相关
	writeToFile(p, "large-pipeline.out")
	printFile("large-pipeline.out")
}
func main() {
	p := createNetworkPipeline("large.in", 32000, 4)

	writeToFile(p, "large-networkpipeline.out")
	printFile("large-networkpipeline.out")
}

//文件拆分为chunCount块，每块chunkSize创建管道
func createPipeline(filename string, filesize, chunkCount int) <-chan int {
	chunkSize := filesize / chunkCount
	pipeline.Init()
	//chan数组 存放排好序的chan
	// []<-chan int是类型，表示一个数组，里面的东西是<-chan int。{}是里面的数据，一开始是空的。
	// 也可以写成var sortResult []<-chan int。这样的话sortResult就是nil，也可以当做空的slice来用。
	var sortResults []<-chan int
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		file.Seek(int64(i*chunkSize), 0)
		//读文件
		source := pipeline.ReaderSource(bufio.NewReader(file), chunkSize)
		sortResults = append(sortResults, pipeline.InMemSort(source))
	}
	//归并，返回所有元素的chan
	return pipeline.MergeN(sortResults...)
}
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipeline.ReaderSource(file, -1)
	count := 0
	for v := range p {
		if count < 100 {
			fmt.Println(v)
			count++
		}
	}
}

func writeToFile(p <-chan int, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
 	//使用bufio自带的缓冲区，减少数据频繁的读写
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	pipeline.WriteSink(writer, p)
}
func createNetworkPipeline(filename string, filesize, chunkCount int) <-chan int {
	chunkSize := filesize / chunkCount
	pipeline.Init()
	var sortAddr []string
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		file.Seek(int64(i*chunkSize), 0)
		//使用bufio自带的缓冲区，减少数据频繁的读写
		source := pipeline.ReaderSource(bufio.NewReader(file), chunkSize) //读取chunkSize的数据
		addr := ":" + strconv.Itoa(7000+i)
		//每个node sort 后 创建了一个server，等待连接
		pipeline.NetworkSink(addr, pipeline.InMemSort(source))
		sortAddr = append(sortAddr, addr)
	}
	//可以手动 telnet localhost 7000(7001,7002...)
	//return nil

	//做了很多client，每个client连接上面的server,下面是merge的过程
	var sortResults []<-chan int
	for _, addr := range sortAddr {
		//连接sever,从某个server读取已经排序好的数据
		sortResults = append(sortResults, pipeline.NetworkSource(addr))
	}
	return pipeline.MergeN(sortResults...)
}
