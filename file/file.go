package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"os"
)

const FILE_NAME = "file.in"
const N = 5

func main() {

	file, e := os.Create(FILE_NAME)
	if e != nil {
		panic(e)
	}
	defer file.Close()

	p := RandomSource(N)
	//channel引用类型，不能遍历chan，这样会排空chan，下面写文件会为空
	/*fmt.Println("产生的随机数如下：")
	for v := range p {
		fmt.Println(v)
	}*/
	fmt.Println("写入文件：")
	writer := bufio.NewWriter(file)
	WriteSink(writer, p)
	writer.Flush()
	file, e = os.Open(FILE_NAME)
	if e != nil {
		panic(e)
	}
	defer file.Close()
	fmt.Println("读文件结果如下：")
	/*p1 := ReaderSource(file, -1)
	for v1 := range p1 {
		fmt.Println(v1)
	}*/
	read := make(chan int, 1024)
	go ReaderSource2(read, file, -1)
	for v1 := range read {
		fmt.Println(v1)
	}
}
func ReaderSource2(out chan<- int, reader io.Reader, chunkSize int) {
	buffer := make([]byte, 8)
	bytesRead := 0
	for {
		n, err := reader.Read(buffer)
		bytesRead += n
		if n > 0 {
			v := int(binary.BigEndian.Uint64(buffer))
			out <- v
		}
		if err != nil || (chunkSize != -1 && bytesRead >= chunkSize) {
			break
		}
	}
	close(out)
}

func ReaderSource(reader io.Reader, chunkSize int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		buffer := make([]byte, 8)
		bytesRead := 0
		for {
			n, err := reader.Read(buffer)
			bytesRead += n
			if n > 0 {
				v := int(binary.BigEndian.Uint64(buffer))
				out <- v
			}
			if err != nil || (chunkSize != -1 && bytesRead >= chunkSize) {
				break
			}
		}
		close(out)
	}()
	return out
}

func WriteSink(writer io.Writer, in <-chan int) {
	for v := range in {
		buffer := make([]byte, 8)
		binary.BigEndian.PutUint64(buffer, uint64(v))
		writer.Write(buffer)
	}
}
func RandomSource(count int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Int()
		}
		close(out)
	}()
	return out
}
