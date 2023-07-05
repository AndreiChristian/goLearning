package tcp

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

type Server struct {
	listener net.Listener
}

func New(adress string) (*Server, error) {

	listener, err := net.Listen("tcp", adress)

	if err != nil {

		return nil, fmt.Errorf("error creating server : %w", err)

	}

	return &Server{listener: listener}, nil

}

func (s *Server) Listen() error {

	for {

		conn, err := s.listener.Accept()

		if err != nil {

			return fmt.Errorf("error accepting connections : %w", err)

		}

		s.handleConnection(conn)

	}

}

func (s *Server) handleConnection(conn net.Conn) {

	defer conn.Close()

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {

		text := scanner.Text()
		io.WriteString(conn, text+"\n")

	}

}
