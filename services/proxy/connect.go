package proxy

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func isCONNECT(content string) bool {
	if content[:7] == "CONNECT" {
		return true
	}
	return false
}

func readCONNECTALLData(conn net.Conn, buffer []byte) ([]byte, error) {
	receiveBuffer := make([]byte, 4096)
	fmt.Fprint(conn, "HTTP/1.1 200 Connection established\r\n"+
		"Connection: keep-alive\r\n"+
		"Via: 1.1 localhost\r\n")
	//conn.Read(receiveBuffer)
	//fmt.Print(receiveBuffer)
	//if err != nil {
	//  return buffer, err
	//}
	return append(buffer, receiveBuffer...), nil
}

func parseCONNECTData(b []byte) *httpData {
	return &httpData{}
}

func (p *Proxy) handleCONNECT(conn net.Conn, buffer []byte) bool {
	content := string(buffer)
	if !isCONNECT(content) {
		return false
	}
	//data, _ := readCONNECTALLData(conn, buffer)
	// parseData := parseCONNECTData(data)
	proxyConn, err := net.DialTimeout("tcp", "192.168.43.1:8080", time.Second)
	if err != nil {
		log.Println(err.Error())
		fmt.Fprint(conn, "HTTP/1.1 500 Internal server error\r\n")
		fmt.Fprint(conn, "Content-Type: text/plain\r\n")
		fmt.Fprint(conn, "Connection: close\r\n")
		fmt.Fprint(conn, "\r\n")
		fmt.Fprint(conn, err.Error())
		fmt.Fprint(conn, "\r\n")
		conn.Close()
		return true
	}
	go func() {
		for {
			receiveBuffer := make([]byte, 1024*64)
			_, err := proxyConn.Read(receiveBuffer)
			if err != nil {
				fmt.Println("cerrado 1", string(receiveBuffer), err)
				conn.Close()
				proxyConn.Close()
				break
			}

			conn.Write(receiveBuffer)
			os.Stdout.Write(receiveBuffer)
		}
	}()
	go func() {
		for {
			buffer := make([]byte, 1024*64)
			_, err := conn.Read(buffer)
			if err != nil {
				fmt.Println("cerrado 2", err)
				conn.Close()
				proxyConn.Close()
				break
			}
			receiveBuffer := []byte(strings.Trim(string(buffer), string([]byte{0})))

			proxyConn.Write(receiveBuffer)
			os.Stdout.Write(receiveBuffer)
		}
	}()
	os.Stdout.Write([]byte("###############################\n"))
	os.Stdout.Write(buffer)
	proxyConn.Write(buffer)
	return true
}
