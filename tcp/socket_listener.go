package tcp

import (
	"net"
	"strings"
	"sync"
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
func (s *SocketListener) Listen() (err error) {
	var (
		tcpAddr *net.TCPAddr
		mutex   sync.Mutex // Mutex must be inited as an instance
	)

	tcpAddr, _ = net.ResolveTCPAddr("tcp", s.Address+":"+s.Port)

	ln, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		logrus.Error("TCP listen err", err)
		return
	}
	defer ln.Close()

	logrus.Info("TCP start listening")

	for {
		conn, _ := ln.AcceptTCP()
		go s.handleRead(conn, &mutex)
	}
}

// KeepAlive try to keep connection alive
func (s *SocketListener) keepAlive(conn *net.TCPConn) error {
	// Continue if no keepAlivePeriod
	if s.KeepAlivePeriod == nil {
		return nil
	}
	// Error when fail to set keep alive
	if err := conn.SetKeepAlive(true); err != nil {
		return err
	}
	return conn.SetKeepAlivePeriod(*s.KeepAlivePeriod)
}

// HandleRead read & process single data
func (s *SocketListener) handleRead(conn *net.TCPConn, mutex *sync.Mutex) {
	mutex.Lock()

	// Close if disconnected
	defer conn.Close()

	defer mutex.Unlock()

	// Get client ip address
	ipStr := conn.RemoteAddr().String()

	// Read message
	for {
		buffer := make([]byte, 1024)
		length, readErr := conn.Read(buffer)

		if readErr != nil {
			logrus.Info("Read message error", readErr)
			break
		}

		if length > 12 {
			data := buffer[:length]
			dataStr := strings.TrimSpace(string(data))

			in := strings.Split(dataStr, " ")

			if in[0] == "99" {
				logrus.Info("Received message from ", ipStr, ", Content: ", dataStr, ", Length: ", length)
			}

			conn.Write(data)
		}

	}
}
