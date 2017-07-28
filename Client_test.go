package main

import
(
  "testing"
  "os/exec"
  "log"
  "net"
	"net/rpc/jsonrpc"
)


func TestStartServer(t *testing.T) {
  cmd := exec.Command("./Server","&")
  err := cmd.Start()
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")
	log.Printf("Command finished with error: %v", err)
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
    t.Fatal("eroare")
  }

}
