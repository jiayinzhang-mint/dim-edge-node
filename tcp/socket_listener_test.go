package tcp

import (
	"net"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

func TestListen(t *testing.T) {
	ssl := &SocketListener{
		Address: "localhost",
		Port:    "8000",
	}

	ssl.Listen()
}

func TestWrite(t *testing.T) {
	// Connect to socket server
	conn, err := net.DialTimeout("tcp", "127.0.0.1:8000", 5*time.Second)
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
	defer conn.Close()
	for i := 0; i < 5; i++ {
		// Send data
		data := "hello socket, this is message " + strconv.Itoa(i)
		conn.Write([]byte(data))
		time.Sleep(1 * time.Second)
	}
}
