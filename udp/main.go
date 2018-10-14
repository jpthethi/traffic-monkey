package udp

import (
	"log"
	"net"
)

//Start a UDP Connection
func Start() *net.UDPConn {
	hostName := "localhost"
	portNum := "6000"
	service := hostName + ":" + portNum
	RemoteAddr, _ := net.ResolveUDPAddr("udp", service)
	conn, _ := net.DialUDP("udp", nil, RemoteAddr)

	log.Printf("Established connection to %s \n", service)
	log.Printf("Remote UDP address : %s \n", conn.RemoteAddr().String())
	log.Printf("Local UDP client address : %s \n", conn.LocalAddr().String())

	return conn

}
