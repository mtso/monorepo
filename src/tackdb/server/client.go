package server

import (
	"bufio"
	"io"
	"net"
	"strings"
)

type client struct {
	conn   net.Conn
	reader io.Reader
	id     float64
	table  CommandTable
}

func NewClient(conn net.Conn, id float64) *client {
	c := &client{
		conn:   conn,
		id:     id,
		reader: bufio.NewReader(conn),
	}
	c.table = newCommandTable(c)
	return c
}

func (c *client) Write(msg string) error {
	return c.conn.Write([]byte(msg))
}

func (c *client) Read() (string, error) {
	return c.reader.ReadString('\n')
}

func (c *client) HandleCommand() error {
	msg, err := c.Read()
	if err == io.EOF {
		return NewErrClientClosed(c.id)
	}

	msg = strings.Trim(msg)

	cmd, err := ParseCommand(msg)
}

func (c *client) ReadArgs() ([]string, error) {
	buf, err := c.reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	buf = strings.Trim(buf, "\n")
	args := strings.Split(buf, " ")
	return args, nil
}

func (c *client) Serve() (err error) {
	defer c.conn.Close()

	c.conn.Write(startupmsg)
	log.Printf("connected. id=%v", c.id)

	for {
		args, err := client.ReadArgs()
		if err != nil {
			log.Println(err)
			break
		}

		cmdname, args := args[0], args[1:]
		cmd, ok := c.table[cmdname]
		if !ok {
			c.conn.Write(ErrUnrecognizedCommand.Error())
			continue
		}

		res, err := cmd(args...)
		if err == ErrConnectionClosed {
			break
		} else if err != nil {
			log.Println(err)
			client.Write(err.Error())
		} else {
			client.Write([]byte(res))
		}
	}
	return
}

func (c *client) newCommandTable(name string) CommandTable {
	return &CommandTable{
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
}
