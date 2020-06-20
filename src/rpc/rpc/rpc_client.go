package main

import (
	"fmt"
	"log"
	"net/rpc"
)

// 运算请求结构体
type AirthRequest struct {
	A int
	B int
}

// 运算响应结构体
type AirthResponse struct {
	Pro int // 乘积
	Quo int // 商
	Rem int // 余数
}

func main() {
	conn, err := rpc.DialHTTP("tcp", "127.0.0.1:8905")
	if err != nil {
		log.Fatalln("dialing error:", err)
	}
	req := AirthRequest{10, 2}
	var res AirthResponse
	err = conn.Call("Arith.Multiply", req, &res)

	if err != nil {
		log.Fatalln("arith error:", err)
	}
	fmt.Println("%d * %d = %d\n", req.A, req.B, res.Pro)
	err = conn.Call("Arith.Divide", req, &res)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("%d / %d, quo is %d, rem is %d\n", req.A, req.B, res.Quo, res.Rem)
}
