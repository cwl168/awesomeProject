package pipeline

import (
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"sort"
	"time"
)

var startTime time.Time

func Init() {
	startTime = time.Now()
}

//<-chan int 对于用的人来说只能拿东西，对于自己而言只能放东西,返回只读chan
func ArraySource(a ...int) <-chan int {
	out := make(chan int)
	// out <- a[0] 不能这样往通道写数据，不能自己给自己导入数据，必须起一个goroutine
	go func() {
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

func InMemSort(in <-chan int) <-chan int {
	//定义为缓冲通道，非无缓存通道，不用阻塞等待，提高了管道处理的效率
	out := make(chan int, 1024)
	go func() {
		//读入内存
		a := []int{}
		for v := range in {
			a = append(a, v)
		}
		fmt.Println("Read done:", time.Now().Sub(startTime))
		//排序
		sort.Ints(a)
		fmt.Println("InMemSort done:", time.Now().Sub(startTime))
		//输出
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}
func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		// 归并的过程中要处理某个通道中可能没有数据的情况
		v1, ok1 := <-in1 //会等待
		v2, ok2 := <-in2
		for ok1 || ok2 {
			//fmt.Printf("v1 %s,v2 %s",v1,v2)
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1
				v1, ok1 = <-in1
			} else {
				out <- v2
				v2, ok2 = <-in2
			}
		}
		close(out)
		fmt.Println("Merge done:", time.Now().Sub(startTime)) //merge是同时进行，合并的次数
	}()
	return out
}

//chunkSize 分块读取,值为-1表示全部读取文件 将reader改为分块读取, 每次读取指定字符长度
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

//递归归并
func MergeN(inputs ...<-chan int) <-chan int {
	if len(inputs) == 1 {
		return inputs[0]
	}
	m := len(inputs) / 2
	//merge inputs[0...m) and inputs[m...end)
	return Merge(
		MergeN(inputs[:m]...),
		MergeN(inputs[m:]...),
	)
}
