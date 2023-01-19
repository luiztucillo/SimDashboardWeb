package main

import (
	"bufio"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8001")
	if err != nil {
		panic(err)
	}

	helloMessage, errC := bufio.NewReader(conn).ReadString('\n')

	if errC != nil {
		panic(errC)
	}

	if helloMessage != "Connected\n" {
		panic("Failed connect")
	}

	c, err1 := conn.Write([]byte("1234qwer\n"))

	if err1 != nil || c == 0 {
		panic(err)
	}

	authed, errA := bufio.NewReader(conn).ReadString('\n')

	if errA != nil {
		panic(errA)
	}

	if authed != "OK\n" {
		panic("Authr error")
	}

	readFile, err := os.Open("/Users/ltucillo/Projetos/Pessoal/SimDashboardWeb/desktop_app/games_data.csv")
	defer readFile.Close()

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		text := fileScanner.Text() + "\n"
		conn.Write([]byte(text))
	}

	conn.Close()
}
