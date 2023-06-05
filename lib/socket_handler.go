package lib

import (
	"fmt"
	"net"
)

type Socket_handler struct {
	Conn *net.UDPConn
	Ip   []byte
	Port int
	Err  error
}

func (s *Socket_handler) SocketClient(method string) error {
	//socket_path := fmt.Sprintf("%s:%s", s.Ip, s.Port)
	//defer s.Conn.Close()
	if method == "udp" {
		s.Conn, s.Err = net.DialUDP("udp", nil, &net.UDPAddr{
			IP:   s.Ip,
			Port: 9090,
		})
		if s.Err != nil {
			return s.Err
		}
	}
	
	return nil
}

func (s *Socket_handler) UdpSend(msg []byte) error {
	_, s.Err = s.Conn.Write(msg)

	if s.Err != nil {
		fmt.Printf("send data fail,err :%v\n", s.Err)
		return s.Err
	}
	return nil
}

func (s *Socket_handler) RevdData() error {
	result := make([]byte, 1024)
	//n, err := conn.Read(buf[:])
	n, remoteAddr, err := s.Conn.ReadFromUDP(result)
	if err != nil {
		fmt.Printf("receive data failed, err: %v\n", err)
		return err
	}
	fmt.Printf("receive from addr: %v  data: %v\n", remoteAddr, string(result[:n]))
	return nil
}
