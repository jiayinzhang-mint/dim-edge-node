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
		Port:    "9000",
	}

	ssl.Listen()
}

func TestWrite(t *testing.T) {
	// Connect to socket server
	conn, err := net.DialTimeout("tcp", "127.0.0.1:9000", 5*time.Second)
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
	defer conn.Close()
	for i := 0; i < 10; i++ {
		// Send data
		data := "99 hello socket, this is message " + strconv.Itoa(i)
		n, err := conn.Write([]byte(data))
		logrus.Info(n, err)
		time.Sleep(3 * time.Microsecond)
	}
}
