// client.go
package main

import (

	"log"
	"net"
	"net/rpc/jsonrpc"

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


func main() {

	client, err := net.Dial("tcp", "127.0.0.1:1230")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	/*
  Test sum of 2 numbers
  Item1,_ := strconv.Atoi(os.Args[1])
  Item2,_ := strconv.Atoi(os.Args[2])
	args := &ArgsSum{Item1,Item2}
	var reply int
	c := jsonrpc.NewClient(client)
	err = c.Call("MyServer.Sum", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Result: %d+%d=%d\n", args.Item1, args.Item2, reply)
  */

  /*
  //Test read from file
  var result int
  Path := "/home/teo/Desktop/work/src/my_project/Client_Server/file.txt"
  args :=&ArgsRead{Path}
  c := jsonrpc.NewClient(client)
	err = c.Call("MyServer.Read", args, &result)
	if err != nil {
		log.Fatal("arith error:", err)
	}
  fmt.Printf("Number from file is %d\n",result)
  */

  //Test write to file
  var reply string
  Item := 7
  Path := "/home/teo/Desktop/work/src/my_project/Client_Server/asd.txt"
  args :=&ArgsWrite{Item,Path}
  c := jsonrpc.NewClient(client)
	err = c.Call("MyServer.Write", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}


}
