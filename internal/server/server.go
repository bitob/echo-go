package server

import (
	"echo-go/internal/utils"
	"net"
)

func StartServer(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		utils.Log.Errorf("Error starting the server: %v", err)
	}
	defer listener.Close()

	utils.Log.Infof("Server is listening on %s", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			utils.Log.Errorf("Error accepting connection: %v", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	utils.Log.Infof("Received connection from %s", conn.RemoteAddr())

	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			utils.Log.Errorf("Error reading from connection: %v", err)
			return
		}

		utils.Log.Infof("Received %d bytes: %s\n", n, string(buffer[:n]))

		_, err = conn.Write([]byte("Echo: "))
		if err != nil {
			utils.Log.Errorf("Error writing to connection: %v", err)
			return
		}
		_, err = conn.Write(buffer[:n])
		if err != nil {
			utils.Log.Errorf("Error writing to connection: %v", err)
			return
		}
	}
}
