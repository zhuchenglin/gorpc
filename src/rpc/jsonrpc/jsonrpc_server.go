package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

// 算数运算结构体
type Arith struct {
}

// 算数运算请求结构体
type ArithRequest struct {
	A int
	B int
}

// 算数运算响应结构体
type ArithResponse struct {
	Pro int  // 乘积
	Quo int	  // 商
	Rem int  // 余数
}

// 乘法运算方式
func (this *Arith) Multiply(req ArithRequest, res *ArithResponse) error  {
	res.Pro = req.A * req.B
	return nil
}

// 除法运算方法
func (this *Arith) Divide (req ArithRequest, res *ArithResponse) error  {
	if req.B == 0 {
		return errors.New("divide by zero")
	}
	res.Quo = req.A / req.B
	res.Rem = req.A % req.B
	return nil
}

func main()  {
	rpc.Register(new (Arith))
	lis, err := net.Listen("tcp", "127.0.0.1:9503")
	if err != nil {
		log.Fatalln("fatal error:", err)
	}
	fmt.Fprintf(os.Stdout, "%s", "start connection\n")

	for {
		conn, err := lis.Accept()
		if err != nil {
			continue
		}
		go func(conn net.Conn) {
			fmt.Fprintf(os.Stdout, "%s", "new client in coming\n")
			jsonrpc.ServeConn(conn)
		}(conn)

	}


}
