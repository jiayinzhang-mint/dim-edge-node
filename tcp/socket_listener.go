package tcp

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

// SocketListener socket listener
type SocketListener struct {
	Address         string         `json:"address"`
	Port            string         `json:"port"`
	MaxConnections  int            `json:"maxConnection"`
	KeepAlivePeriod *time.Duration `json:"keepAlivePeriod"`
}

// Listen start socket for async listening
func (s *SocketListener) Listen() {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", s.Address+":"+s.Port)

	ln, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		logrus.Error("TCP listen err")
		os.Exit(1)
	}
	defer ln.Close()

	logrus.Info("TCP start listening")

	for {
		conn, _ := ln.AcceptTCP()
		go s.read(conn)
	}
}

// KeepAlive try to keep connection alive
func (s *SocketListener) keepAlive(conn *net.TCPConn) error {
	// Continue if no keepAlivePeriod
	if s.KeepAlivePeriod == nil {
		return nil
	}
	// Error when set keep alive
	if err := conn.SetKeepAlive(true); err != nil {
		return err
	}
	return conn.SetKeepAlivePeriod(*s.KeepAlivePeriod)
}

// Read ready for processing single data
func (s *SocketListener) read(conn *net.TCPConn) {
	ipStr := conn.RemoteAddr().String()

	// Close if disconnected
	defer func() {
		logrus.Info("TCP disconnected :" + ipStr)
		conn.Close()
	}()

	// Read message
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		fmt.Println(message)
		msg := time.Now().String() + "\n"
		b := []byte(msg)
		conn.Write(b)
	}
}
