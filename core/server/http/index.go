package http

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

// Server represents the HTTP server
type Server struct {
	addr string
}

// NewServer creates a new HTTP server
func NewServer(addr string) *Server {
	return &Server{addr: addr}
}

// Start runs the server
func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	defer listener.Close()

	log.Printf("Server listening on %s", s.addr)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}
		go s.handleConnection(conn)
	}
}

// handleConnection processes incoming connections
func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	// Read request line
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Error reading request: %v", err)
		return
	}

	// Parse request line
	parts := strings.Fields(line)
	if len(parts) < 3 {
		s.writeResponse(conn, 400, "Bad Request", "Invalid request")
		return
	}

	method, path := parts[0], parts[1]
	if method != "GET" {
		s.writeResponse(conn, 405, "Method Not Allowed", "Only GET is supported")
		return
	}

	// Simple routing
	body := "Hello from Go HTTP Server!"
	if path != "/" {
		s.writeResponse(conn, 404, "Not Found", "Page not found")
		return
	}

	s.writeResponse(conn, 200, "OK", body)
}

// writeResponse sends an HTTP response
func (s *Server) writeResponse(conn net.Conn, statusCode int, statusText, body string) {
	fmt.Fprintf(conn, "HTTP/1.1 %d %s\r\n", statusCode, statusText)
	fmt.Fprintf(conn, "Content-Type: text/plain\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Server: CustomGoServer/1.0\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, "%s", body)
}