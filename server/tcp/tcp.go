package tcp

import (
	"bufio"
	"fmt"
	"net"
)

func CreateServer(c chan<- string) {
	l, err := net.Listen("tcp4", ":8001")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleTcpConnection(conn, c)
	}
}

func closeConnection(conn net.Conn) {
	fmt.Printf("Close %s\n", conn.RemoteAddr().String())
	conn.Close()
}

func handleTcpConnection(conn net.Conn, c chan<- string) {
	fmt.Printf("Serving %s\n", conn.RemoteAddr().String())

	conn.Write([]byte("Connected\n"))

	auth, err := bufio.NewReader(conn).ReadString('\n')

	if err != nil {
		conn.Write([]byte("Auth Error\n"))
		closeConnection(conn)
		return
	}

	if auth != "1234qwer\n" {
		conn.Write([]byte("Auth Invalid\n"))
		closeConnection(conn)
		return
	} else {
		conn.Write([]byte("OK\n"))
	}

	scanner := bufio.NewScanner(conn)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		c <- scanner.Text() + "\n"
	}
}
