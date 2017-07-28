// client.go
package main

import (

	"log"
	"net"
	"net/rpc/jsonrpc"
	"fmt"

)

type ArgsSum struct {
    Item1, Item2 int
}

type ArgsRead struct {
    FilePath string
}


type ArgsWrite struct {
		Item int
    FilePath string
}

func Plus(x int , y int) int{
	client, err := net.Dial("tcp", "127.0.0.1:1232")
  if err != nil {
    log.Fatal("dialing:", err)
  }

  //Test sum of 2 numbers
  Item1 := x
  Item2 := y
	args := &ArgsSum{Item1,Item2}
	var reply int
	c := jsonrpc.NewClient(client)
	err = c.Call("MyServer.Sum", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	return reply
}

func Read() int{
	client, err := net.Dial("tcp", "127.0.0.1:1232")
  if err != nil {
    log.Fatal("dialing:", err)
  }
	var result int
  Path := "./file.txt"
  args :=&ArgsRead{Path}
  c := jsonrpc.NewClient(client)
	err = c.Call("MyServer.Read", args, &result)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	return result
}

func Write() int{
	client, err := net.Dial("tcp", "127.0.0.1:1232")
  if err != nil {
    log.Fatal("dialing:", err)
  }
	var result int
	Item := 15
  Path := "./asder.txt"
  args :=&ArgsWrite{Item,Path}
  c := jsonrpc.NewClient(client)
	err = c.Call("MyServer.Write", args, &result)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	return result
}

func main() {
	fmt.Println(Write())
}
