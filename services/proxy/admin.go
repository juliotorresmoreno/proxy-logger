package proxy

import "bytes"

type connection struct {
	buffer *bytes.Buffer
}

func newConnection() *connection {
	buff := make([]byte, 0)
	return &connection{
		buffer: bytes.NewBuffer(buff),
	}
}

// AdminServer Administrador de conexiones
type AdminServer struct {
	connection      chan *connection
	closeConnection chan *connection
	connections     map[*connection]bool
	close           chan bool
}

func newAdminServer() *AdminServer {
	return &AdminServer{
		closeConnection: make(chan *connection),
		connections:     make(map[*connection]bool),
		connection:      make(chan *connection),
		close:           make(chan bool),
	}
}

// Listen Habilita un socket para escuchar peticiones
func (a *AdminServer) Listen() {
	for {
		select {
		case connection := <-a.connection:
			a.connections[connection] = true
		case connection := <-a.closeConnection:
			delete(a.connections, connection)
		case <-a.close:
			break
		}
	}
}
