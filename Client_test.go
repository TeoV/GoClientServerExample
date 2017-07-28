package main

import
(
  "testing"
  "os/exec"
  "log"
  "net"
	"net/rpc/jsonrpc"
  "time"
)


func TestStartServer(t *testing.T) {
  time.Sleep(100*time.Millisecond)
  cmd := exec.Command("./Server","&")
  err := cmd.Start()
	if err != nil {
		t.Fatal(err)
	}
}


func TestPlus(t *testing.T) {
  client, err := net.Dial("tcp", "127.0.0.1:1232")
  if err != nil {
    log.Fatal("dialing:", err)
  }
  Item1 := 1
  Item2 := 2
	args := &ArgsSum{Item1,Item2}
	var reply int
	c := jsonrpc.NewClient(client)
	err = c.Call("MyServer.Sum", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
  if reply != 3 {
    t.Fatal("Invalid sum")
  }

}

func TestRead(t *testing.T){
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
  if result !=123 {
    t.Fatal("Invalid number")
  }
}

func TestWrite(t *testing.T){
  client, err := net.Dial("tcp", "127.0.0.1:1232")
  if err != nil {
    log.Fatal("dialing:", err)
  }
	var result int
  Item := 12
  Path := "./asder.txt"
  args :=&ArgsWrite{Item,Path}
  c := jsonrpc.NewClient(client)
	err = c.Call("MyServer.Write", args, &result)
	if err != nil {
		log.Fatal("arith error:", err)
	}

  if result != Item {
  t.Fatal("Error at Writing")
}
}
