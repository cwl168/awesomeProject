package mergeSort

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"os"
)

func ReaderSource(reader io.Reader) <-chan int {
	out := make(chan int)
	go func() {
		// 这里定义为8个字节, 原因是我的机器是64位的, 所以int也是64位, 那么对应的字节数就是8个字节
		buffer := make([]byte, 8)
		for {
			// reader返回两个参数, 第一个是读取到的字节数, 第二个是err异常
			n, err := reader.Read(buffer)
			if n > 0 {
				// 如果读到了, 就把读到的东西发给channel
				u := binary.BigEndian.Uint64(buffer)
				out <- int(u)
			}

			if err != nil {
				break
			}
		}
		close(out)
	}()
	return out
}
func WriteSink(writer io.Writer, in <-chan int) {
	for v := range in {
		b := make([]byte, 8)
		binary.BigEndian.PutUint64(b, uint64(v))
		writer.Write(b)
	}
}
func RandomSource(count int) chan int {
	out := make(chan int)
	go func() {
		// 生成count个随机数
		for i := 0; i < count; i++ {
			out <- rand.Int()
		}

		close(out)
	}()
	return out
}

func Gen() {
	// 第一步: 造数据, 生成100个随机数, 写入到文件
	const fileName = "file_sort.txt"
	const count = 100
	// 第一步: 将随机生产的数字保存到small.in文件
	// 构造第一个数据源
	file, e := os.Create(fileName)
	if e != nil {
		panic(e)
	}
	defer file.Close()
	//产生随机数
	dataSource := RandomSource(count)

	//往文件写
	writer := bufio.NewWriter(file)
	WriteSink(writer, dataSource)
	writer.Flush()

	// 第二步: 从文件中读取文件内容, 在控制台打印
	// 从第一个数据源读取出数据
	f, e := os.Open(fileName)
	if e != nil {
		panic(e)
	}
	defer f.Close()
	readerSource := ReaderSource(bufio.NewReader(f))

	var num = 0
	for rs := range readerSource {
		fmt.Println(rs)
		num++

	}
}
