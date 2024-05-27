package wsHandler

import (
	"fmt"
	"io"
	"log"

	"github.com/google/uuid"
	"golang.org/x/net/websocket"
)

type client struct {
	ID   string
	name string
	conn *websocket.Conn
}

type message struct {
	senderID string
	content  string
}

type Server struct {
	conns map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) HandleConn(ws *websocket.Conn) {

	fmt.Println("New connection from client", ws.RemoteAddr())
	client := client{
		ID:   uuid.NewString(),
		conn: ws,
	}

	fmt.Println("Client has been assigned with ID: ", client.ID)

	s.conns[ws] = true

	// TODO: Implement Mutex
	s.readLoop(ws)
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)

	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("read error: ", err)
			continue
		}

		msg := buf[:n]
		log.Println("message received", string(msg))

		s.broadcast(msg)
	}
}

func (s *Server) broadcast(b []byte) {
	for ws := range s.conns {
		go func(ws *websocket.Conn) {
			_, err := ws.Write(b)

			if err != nil {
				fmt.Println("write error: ", err)
			}
		}(ws)
	}
}
