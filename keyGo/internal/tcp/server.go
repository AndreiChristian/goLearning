package tcp

import (
	"bufio"
	"fmt"
	"io"
	"keygo/internal/protocol"
	"keygo/pkg/store"
	"net"
)

type Server struct {
	listener net.Listener
	store    *store.Store
}

func New(adress string) (*Server, error) {

	listener, err := net.Listen("tcp", adress)

	if err != nil {

		return nil, fmt.Errorf("error creating server : %w", err)

	}

	return &Server{listener: listener, store: store.New()}, nil

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
		command := protocol.ParseCommand(scanner.Text())
		result := s.executeCommand(command)
		io.WriteString(conn, result+"\n")

	}

}

func (s *Server) executeCommand(command protocol.Command) string {

	switch command.Name {

	case "SET":

		if len(command.Args) != 2 {
			return "ERR wrong number of arguments for 'set' command"
		}

		s.store.Set(command.Args[0], command.Args[1])
		return "ok"

	case "GET":

		if len(command.Args) != 1 {

			return "ERR wrong number of arguments for 'get' command"

		}

		value, ok := s.store.Get(command.Args[0])

		if !ok {
			return "(nil)"
		}

		return value

	default:
		return fmt.Sprintf("ERR unknown command '%s'", command.Name)
	}

}
