package main

import (
	"encoding/json"
	"fmt"
	"net"
	"testgo/socket/go_socket/lib"
	"time"
)

func main() {

	sc := lib.Socket_handler{
		Ip:   net.IPv4(127, 0, 0, 1),
		Port: 9090,
	}
	sc.SocketClient("udp")
	data := make(map[string]interface{})
	data["method"] = "screw"
	data["state"] = "2"
	js_l, _ := json.Marshal(data)
	sc.UdpSend(js_l)
	go func() {
		for {
			sc.RevdData()
			fmt.Println("receive data")
			time.Sleep(2 * time.Second)
			//os.Exit(1)
		}
	}()
	time.Sleep(3 * time.Second)
}
