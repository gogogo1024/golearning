package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/rpc"
)


type Args struct {
	A , B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}


func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

// 注册服务对象并开启RPC服务
func main()  {
    arith :=new(Arith)
    rpc.Register(arith)
    rpc.HandleHTTP()
    if err :=http.ListenAndServe(":1234",nil) ; err != nil {
		log.Fatal("listen error",err)
	}
	fmt.Println("server listen at port 1234 successfully")
}


