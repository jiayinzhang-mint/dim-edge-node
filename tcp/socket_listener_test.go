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
	conn, err := net.DialTimeout("tcp", "192.168.64.14:31986", 5*time.Second)
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
	defer conn.Close()
	i := 0
	for {
		// Send data
		data := "99 hello socket, this is message " + strconv.Itoa(i) + "from 1"
		i++
		n, err := conn.Write([]byte(data))
		logrus.Info(n, err)
		time.Sleep(100000 * time.Microsecond)
	}
}
