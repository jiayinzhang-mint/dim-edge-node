package tcp

import "testing"

func TestListen(t *testing.T) {
	ssl := &SocketListener{
		Address: "localhost",
		Port:    "8000",
	}

	ssl.Listen()
}
