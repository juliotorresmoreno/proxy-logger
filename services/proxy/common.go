package proxy

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

type validConn struct {
	uuid string
	net.Conn
	status bool
}

func (el validConn) Write(b []byte) (int, error) {
	os.Stdout.Write(b)
	return el.Conn.Write(b)
}

func (el validConn) Close() error {
	return el.Conn.Close()
}

func pipe(src, dst *validConn) {
	defer func() {
		recover()
		src.Close()
	}()
	for {
		if !src.status {
			break
		}
		src.SetReadDeadline(time.Now().Add(time.Second * 3600))
		buffer, err := getUpgrade(src)
		if err == io.EOF {
			dst.status = false
			break
		}
		if len(buffer) > 0 {
			_, err = dst.Write(buffer)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func getProxyHost() string {
	return "127.0.0.1:6000"
}
