package server

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

const (
	SCHEME  = "tcp"
	VERSION = "0.0.1"
)

var ErrQuit = errors.New("SERVER QUIT")
var ErrUnrecognizedCommand = errors.New("UNRECOGNIZED COMMAND")
var ErrNil = errors.New("NULL")
var ErrNoKey = errors.New("NO KEY")
var ErrNoValue = errors.New("NO VALUE")

var startupmsg = []byte("tackdb v" + VERSION + "\n")

type Server struct {
	conn *net.Listener
	port string
}

func (s *Server) Close() error {
	return (*s.conn).Close()
}

func (s *Server) Start(port string) error {
	s.port = port
	listener, err := net.Listen(SCHEME, ":"+port)
	if err != nil {
		return err
	}
	s.conn = &listener
	InitStore()

	return s.handle()
}

func (s *Server) handle() error {
	clientid := 1.0
	for {
		conn, err := (*s.conn).Accept()
		if err != nil {
			return err
		}

		client := NewClient(conn, clientid)

		go client.Serve()
		clientid += 1

		// go handleClient(client, clientid)
		// clientid += 1
	}
}

var ErrClientClosed = errors.New("client closed")

func NewErrClientClosed(id float64) error {
	errmsg := fmt.Sprintf("client closed. %v", id)
	return errors.New(errmsg)
}

func parseCommand(r *io.Reader) error {
	msg, err := r.ReadString('\n')
	if err == io.EOF {
		log.Printf("client closed. id=%v", id)
		return NewErrClientClosed()
	} else if err != nil {
		log.Println(err)
		continue
	}

	// msg := string(bmsg)
	msg = strings.Trim(msg, "\n")
	log.Printf("%q", msg)

	args := strings.Split(msg, " ")

	// Guard against an empty string.
	if len(args) == 1 && args[0] == "" {
		log.Printf("SKIP command=%q", args)
	}

	command := strings.ToUpper(args[0])
	cmd, ok := commands[command]

	if !ok {
		conn.Write([]byte(ErrUnrecognizedCommand.Error() + "\n"))
		continue
	}

	res, err := cmd(args[1:]...)
	log.Printf("id=%v %s %s", id, res, err)
	if err != nil {
		conn.Write([]byte(err.Error() + "\n"))
	} else {
		conn.Write([]byte(res + "\n"))
	}
}

func handleClient(conn net.Conn, id float64) {
	conn.Write(startupmsg)
	log.Printf("client connected. id=%v", id)

	reader := bufio.NewReader(conn)

	for {
		msg, err := reader.ReadString('\n')
		if err == io.EOF {
			log.Printf("client closed. id=%v", id)
			break
		} else if err != nil {
			log.Println(err)
			continue
		}

		// msg := string(bmsg)
		msg = strings.Trim(msg, "\n")
		log.Printf("%q", msg)

		args := strings.Split(msg, " ")

		// Guard against an empty string.
		if len(args) == 1 && args[0] == "" {
			log.Printf("SKIP command=%q", args)
		}

		command := strings.ToUpper(args[0])
		cmd, ok := commands[command]

		if !ok {
			conn.Write([]byte(ErrUnrecognizedCommand.Error() + "\n"))
			continue
		}

		res, err := cmd(args[1:]...)
		log.Printf("id=%v %s %s", id, res, err)
		if err != nil {
			conn.Write([]byte(err.Error() + "\n"))
		} else {
			conn.Write([]byte(res + "\n"))
		}
	}
}

func trim(str string) string {
	if str[len(str)-2:len(str)-1] == "\n" {
		return str[:len(str)-1]
	}
	return str
}
