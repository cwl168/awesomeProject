package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type Human struct {
	Name string
	Age  int
}

func Serialize(human *Human) []byte {
	// 声明一个字节缓冲区
	var buffer bytes.Buffer
	// 构造一个编码器
	encoder := gob.NewEncoder(&buffer)
	// 进行序列化
	err := encoder.Encode(human)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return buffer.Bytes()
}

func Deserialize(buffer []byte) *Human {
	// 声明一个结构体
	var human Human
	// 声明一个字节缓冲区
	// 构造一个解码器
	decoder := gob.NewDecoder(bytes.NewBuffer(buffer))
	// 进行反序列化
	err := decoder.Decode(&human)
	if err != nil {
		//log.Println("ERROR!")
		//return nil
	}
	return &human
}

func main() {
	// 序列化
	human1 := Human{Name: "Jack", Age: 20}
	fmt.Println("开始进行序列化!")
	result := Serialize(&human1)
	fmt.Println(result)
	// 反序列化
	fmt.Println("开始进行反序列化!")
	human2 := Deserialize(result)
	fmt.Println(*human2)
}
