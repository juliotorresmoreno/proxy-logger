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
	fmt.Println(b)
	return el.Conn.Write(b)
}

func (el validConn) Close() error {
	fmt.Println("Conexion cerrada")
	return el.Conn.Close()
}

func pipe(src, dst *validConn) {
	defer func() {
		recover()
		src.Close()
	}()
	cont := 0
	for {
		cont++
		fmt.Println("Ciclo", src.uuid)
		if !src.status {
			break
		}
		src.SetReadDeadline(time.Now().Add(time.Second * 3600))
		buffer := make([]byte, 1024*64)
		n, err := src.Read(buffer)
		if err == io.EOF {
			dst.status = false
			dst.Close()
			break
		} else if err != nil {
			fmt.Println("Error", src.uuid, err.Error())
		}
		if n > 0 {
			_, err = dst.Write(buffer[:n])
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
