package status

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
var subscriptores = map[*websocket.Conn]bool{}

func send(ws *websocket.Conn, message string) {
	err := ws.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		log.Println(err.Error())
	}
}

func Send(message string) {
	for ws := range subscriptores {
		send(ws, message)
	}
}

func listen(ws *websocket.Conn) {
	ticker := time.NewTicker(60 * time.Second)
	defer func() {
		ticker.Stop()
		delete(subscriptores, ws)
		ws.Close()
		fmt.Println("conexión cerrada!")
	}()
	subscriptores[ws] = true
	msgRequests <- true
	ws.SetPongHandler(func(string) error {
		ws.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})
	for {
		tmp := map[string]interface{}{}
		_, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		json.Unmarshal(message, &tmp)
		switch fmt.Sprintf("%v", tmp["type"]) {
		case "get_request":
			fmt.Println("aca")
			if stmp := fmt.Sprintf("%v", tmp["request"]); stmp != "" {
				msgRequestID <- fmt.Sprintf("%v", stmp)
			}
		}
	}
}

func handleGetWebsocket(c echo.Context) error {
	headers := http.Header{}
	headers.Set("Sec-Websocket-Protocol", "echo-protocol")
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), headers)
	if err != nil {
		return err
	}
	go listen(ws)
	return nil
}
