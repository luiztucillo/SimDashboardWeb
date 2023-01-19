package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	readFile, err := os.Open("/Users/ltucillo/Projetos/Pessoal/SimDashboardWeb/streamer/games_data.csv")
	defer readFile.Close()

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
		time.Sleep(time.Second)
	}
}
