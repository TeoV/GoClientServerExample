// server.go
package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
  "fmt"
  "bufio"
  "strconv"
  "os"
)

/*
type ArgsSum struct {
    Item1, Item2 int
}

type ArgsWrite struct {
    Item int
    FilePath string
}

type ArgsRead struct {
    FilePath string
}
*/

type MyServer struct{}

func (srv *MyServer) Sum(args *ArgsSum, reply *int) error {
	*reply = args.Item1 + args.Item2
	return nil
}

 func (srv *MyServer) Write(args *ArgsWrite , reply *bool) bool {
	 file,err := os.Create(args.FilePath)
   if err != nil {
     fmt.Println(err)
   }
   defer file.Close()
    w := bufio.NewWriter(file)
    fmt.Fprintln(w,args.Item)
    w.Flush()

  return true
}

 func (srv *MyServer) Read(args *ArgsRead, reply *int) error {
   file,err := os.Open(args.FilePath)
   if err != nil {
     fmt.Println(err)
   }
   defer file.Close()
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lineStr := scanner.Text()
    num, _ := strconv.Atoi(lineStr)
    *reply = num
}

  return nil
}

func StartServer() string {
	log.Printf("Server start\n")
	cal := new(MyServer)
	server := rpc.NewServer()
	server.Register(cal)
	server.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)
	listener, e := net.Listen("tcp", ":1232")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	test := "succes"

	for {

		log.Printf("Preparing to accept\n")
		conn, err := listener.Accept()
		log.Printf("Accepted\n")
		if err != nil {
			log.Fatal("accept error: " + err.Error())
		}

			log.Printf("Connection established\n")
			go server.ServeCodec(jsonrpc.NewServerCodec(conn))
			defer conn.Close()
			log.Printf("Connection close\n")

}
return test
}
