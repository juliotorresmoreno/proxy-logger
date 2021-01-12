package proxy

import (
	"fmt"
	"io"
	"net"
	"time"
)

type validConn struct {
	net.Conn
	status bool
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
		src.SetReadDeadline(time.Now().Add(time.Second * 30))
		buffer := make([]byte, 1024*64)
		n, err := src.Read(buffer)
		if err == io.EOF {
			dst.status = false
			dst.Close()
			break
		}
		if n > 0 {
			_, err := dst.Write(buffer[:n])
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
