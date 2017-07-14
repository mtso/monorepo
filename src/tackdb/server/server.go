package server

import (
	"bufio"
	"errors"
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
	commands = CommandTable{
		"GET": func(key ...string) (string, error) {
			if len(key) < 1 {
				return "", ErrNoKey
			}
			value, ok := store.Get(key[0])
			if !ok {
				return "", ErrNil
			}
			return value, nil
		},
		"SET": func(args ...string) (string, error) {
			if len(args) < 1 {
				return "", ErrNoKey
			} else if len(args) < 2 {
				return "", ErrNoValue
			}
			key := args[0]
			value := args[1]
			store.Set(key, value)
			return "SET 1", nil
		},
	}
	return s.handle()
}

func (s *Server) handle() error {
	clientid := 1.0
	for {
		client, err := (*s.conn).Accept()
		if err != nil {
			return err
		}
		go handleClient(client, clientid)
		clientid += 1
	}
}

var commands CommandTable

type CommandTable map[string]Command

type Command func(...string) (string, error)

func handleClient(conn net.Conn, id float64) {
	conn.Write(startupmsg)

	reader := bufio.NewReader(conn)

	for {
		bmsg, err := reader.ReadBytes('\n')
		if err == io.EOF {
			log.Printf("client id=%v %s", id, err)
			break
		} else if err != nil {
			log.Println(err)
			continue
		}

		log.Println("got", bmsg)

		msg := string(bmsg)

		args := strings.Split(msg, " ")
		// if len(args) == 1 && args[0][0] == 10 {
		// 	resp := []byte("NOCOMMAND\n")
		// 	log.Printf("Message received (%d): %q: %b\n", len(msg), msg, msg)
		// 	log.Println("unrecognized command")
		// 	conn.Write(resp)
		// 	continue
		// }

		log.Printf("%q", args)
		command := strings.ToUpper(args[0])
		cmd, ok := commands[command]
		log.Println("ok:", ok)
		log.Printf("%q", cmd)
		log.Printf("%q", commands)
		if !ok {
			conn.Write([]byte(ErrUnrecognizedCommand.Error() + "\n"))
			continue
		}

		res, err := cmd(args[1:]...)
		if err != nil {
			conn.Write([]byte(err.Error() + "\n"))
		} else {
			conn.Write([]byte(res + "\n"))
		}

		// log.Println("Received:", msg)
	}
}

func trim(str string) string {
	if str[len(str)-2:len(str)-1] == "\n" {
		return str[:len(str)-1]
	}
	return str
}
