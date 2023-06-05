package service

import (
	"encoding/json"
	"fmt"
	"net"
	"sync"
	"testgo/socket/go_socket/lib"
	"time"
)

type ip_data struct {
	ip   net.IP
	port int
}

var (
	Ips   []ip_data
	Slist = make([]*lib.Socket_handler, 2)
	wg    sync.WaitGroup
)

func IpsConf() []ip_data {
	Ips = append(Ips, ip_data{
		net.IPv4(192, 168, 1, 5),
		9090,
	})
	Ips = append(Ips, ip_data{
		net.IPv4(127, 0, 0, 1),
		9090,
	})
	return Ips
}
func ConnList() []*lib.Socket_handler {
	ip_list := IpsConf()
	fmt.Println(ip_list[0])
	for i, v := range ip_list {
		sc := &lib.Socket_handler{
			Ip:   v.ip,
			Port: v.port,
		}
		sc.SocketClient("udp")
		fmt.Println(sc)
		Slist[i] = sc
	}
	return Slist
}

func M1Ctrl() {
	socket_list := ConnList()

	wg.Add(len(socket_list))
	for i, s := range socket_list {
		go func(s *lib.Socket_handler) {
			fmt.Println("xxxxxxxxxx", i)
			data := make(map[string]interface{})
			data["method"] = "screw"
			data["state"] = "2"
			data["time"] = time.Now()
			js_l, _ := json.Marshal(data)
			s.UdpSend(js_l)
			wg.Done()
		}(s)
	}
	wg.Wait()
}
