package tcp

import (
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
		go s.handleRead(conn)
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

// HandleRead read & process single data
func (s *SocketListener) handleRead(conn *net.TCPConn) {
	ipStr := conn.RemoteAddr().String()

	// Close if disconnected
	defer func() {
		logrus.Info("TCP disconnected :" + ipStr)
		conn.Close()
	}()

	// Read message
	for {
		buffer := make([]byte, 1024)
		length, _ := conn.Read(buffer)

		if length > 12 {
			data := buffer[:length]

			logrus.Info("Received message: ", string(data))
			conn.Write(data)
		}
	}
}
