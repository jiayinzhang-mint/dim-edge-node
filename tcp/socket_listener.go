package tcp

import (
	"net"
	"strings"
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
		// mutex   sync.Mutex  // Mutex must be inited as an instance
	)

	tcpAddr, _ = net.ResolveTCPAddr("tcp", s.Address+":"+s.Port)

	ln, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		logrus.Error("TCP listen err", err)
		return
	}
	defer ln.Close()

	logrus.Info("dim-edge-node TCP start listening at ", s.Address, ":", s.Port)

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
	// Error when fail to set keep alive
	if err := conn.SetKeepAlive(true); err != nil {
		return err
	}
	return conn.SetKeepAlivePeriod(*s.KeepAlivePeriod)
}

// HandleRead read & process single data
func (s *SocketListener) handleRead(conn *net.TCPConn) {
	start := time.Now()

	// Close if disconnected
	defer conn.Close()

	// Get client ip address
	ipStr := conn.RemoteAddr().String()

	// Read message
	for {
		buffer := make([]byte, 1024)
		length, readErr := conn.Read(buffer)

		if readErr != nil {
			logrus.Info(time.Now().Sub(start), " Read message error ", readErr)
			break
		}

		data := buffer[:length]

		if strings.HasPrefix(string(data), "99") {
			logrus.Info("Received message from ", ipStr, ", Content: ", string(data), ", Length: ", length)
		}

		conn.Write(data)

	}
}
