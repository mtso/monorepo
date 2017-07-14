package server

import (
	// "bufio"
	"errors"
	// "net"
	"testing"
	"time"
)

var ErrTimeout = errors.New("timed out")

func TestServer(t *testing.T) {
	errchan := make(chan error)
	server := &Server{}
	go func() {
		errchan <- server.Start("3750")
	}()

	go func() {
		time.Sleep(2000)
		errchan <- nil
	}()

	err := <-errchan
	if err != nil {
		t.Error(err)
	}
	server.Close()

	err = <-errchan
	if err != ErrQuit {
		t.Errorf("Expected %s, but got %s", ErrQuit, err)
	}
}

// func TestMap(t *testing.T) {
// 	port := "3750"
// 	timeout := time.NewTimer(1000)
// 	want := "bar"
// 	start := make(chan error)
// 	done := make(chan error)
// 	close := make(chan error)
// 	server := &Server{}

// 	go func() {
// 		select {
// 		case <-timeout.C:
// 			start <- ErrTimeout
// 			close <- ErrTimeout
// 		}
// 	}()

// 	go func() {
// 		err := server.Start(port)
// 		if err != nil {
// 			start <- err
// 			// t.Fatal(err)
// 		}
// 		defer server.Close()

// 		start <- nil

// 		done <- <-close
// 	}()

// 	err := <-start
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// go func() {
// 		client1, err := net.Dial("tcp", "127.0.0.1:"+port)
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		defer client1.Close()

// 		client1.Write([]byte("SET foo " + want))
// 	// }()

// 	// go func() {
// 		client2, err := net.Dial("tcp", "127.0.0.1:"+port)
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		defer client2.Close()

// 		time.Sleep(100)
// 		client2.Write([]byte("GET foo"))

// 		msg, err := bufio.NewReader(client2).ReadString('\n')
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		if msg != want {
// 			t.Errorf("Expected the response of \"GET foo\" == %s, but got %s", want, msg)
// 		}
// 		close <- nil
// 	// }()

// 	err = <-done
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	isStopped := timeout.Stop()
// 	if !isStopped {
// 		t.Errorf("timeout should not have expired.")
// 	}
// }
