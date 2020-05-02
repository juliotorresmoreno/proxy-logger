package server

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

func isGETRequest(b []byte) bool {
	head := string(b)
	if head[:3] != "GET" {
		return false
	}
	if !strings.Contains(head, "HTTP/1.1") {
		return false
	}
	return true
}
func isCONNECTRequest(b []byte) bool {
	head := string(b)
	if head[:7] != "CONNECT" {
		return false
	}
	if !strings.Contains(head, "HTTP/1.1") {
		return false
	}
	return true
}

func getBodyLen(b []byte) int {
	head := strings.ToLower(string(b))
	pos := strings.Index(head, "content-length")
	defaultValue := 256 * 1024
	if pos == -1 {
		return defaultValue
	}
	tmp := head[pos+15:]
	pos = strings.Index(tmp, "\n")
	tmp = strings.TrimSpace(tmp[:pos])
	if tmp != "" {
		length, err := strconv.Atoi(tmp)
		if err != nil {
			return defaultValue
		}
		return length
	}
	return defaultValue
}

func (svr Server) getUpgrade(conn net.Conn) ([]byte, error) {
	var b []byte
	head := make([]byte, 1024*8)
	n, err := conn.Read(head)
	if err != nil {
		return b, err
	}
	head = head[:n]
	if isTCP(head) {
		return head, nil
	}
	if isGETRequest(head) {
		return head, nil
	}
	if isCONNECTRequest(head) {
		return head, nil
	}
	pos := len(head) - 4
	if pos <= 0 || string(head[len(head)-4:]) != sep+sep {
		return head, nil
	}
	bodyLength := getBodyLen(head) + 1
	body := make([]byte, bodyLength)
	n, err = conn.Read(body)
	if err != nil {
		return b, err
	}
	body = body[:n]
	b = []byte(string(head) + string(body))
	return b, nil
}

func (svr Server) handle(conn net.Conn) error {
	defer func() {
		/*if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			conn.Close()
		}*/
	}()
	b, err := svr.getUpgrade(conn)
	if err != nil {
		fmt.Println(err)
		return err
	}
	request := parseHTTP(b)

	if request.isCONNECT {
		return svr.handleCONNECT(conn, b)
	}
	if request.isTCP {
		return svr.handleTCP(conn, b, false, nil)
	}

	defer conn.Close()
	t, err := svr.handleHTTP(request)
	if err != nil {
		return err
	}
	_, err = conn.Write(t)
	return err
}
