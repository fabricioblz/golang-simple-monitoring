package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoringTimes = 2
const delay = 5

func showIntro() {
	var name string = "World"
	var version float32 = 1.0
	fmt.Println("Hello", name, "!")
	fmt.Println("Version:", version)
}

func showOptions() {
	fmt.Println("1- Start monitoring")
	fmt.Println("2- Logs")
	fmt.Println("0- Exit")
}

func readOption() int {
	var option int
	fmt.Scan(&option)
	return option
}

func readFromSiteFile() []string {
	var sites []string
	file, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fileReader := bufio.NewReader(file)
	for {
		line, err := fileReader.ReadString('\n')
		line = strings.TrimSpace(line)
		fmt.Println(line)
		sites = append(sites, line)
		if err == io.EOF {
			break
		}
	}
	file.Close()
	return sites
}

func monitoring() {
	fmt.Println("Monitoring...")
	sites := readFromSiteFile()
	// sites := []string{"https://random-status-code.herokuapp.com", "https://www.alura.com.br", "https://www.caelum.com.br", "https://www.google.com"}
	// sites = append(sites, "https://www.youtube.com")

	for i := 0; i < monitoringTimes; i++ {
		for _, site := range sites {
			testingSite(site)
		}
		time.Sleep(delay * time.Second)
	}
}

func logMonitoring(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error:", err)
	}
	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + fmt.Sprint(status) + "\n")
	file.Close()

}

func showLogs() {
	file, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(string(file))
}

func showLogsUsingOs() {
	file, err := os.Open("log.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fileReader := bufio.NewReader(file)
	for {
		line, err := fileReader.ReadString('\n')
		fmt.Println(line)
		if err == io.EOF {
			break
		}
	}
	file.Close()
}

func testingSite(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Error:", err)
	}
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "was loaded successfully!")
		logMonitoring(site, true)
	} else {
		fmt.Println("Site:", site, "was loaded with error:", resp.StatusCode)
		logMonitoring(site, false)
	}
}

func main() {
	showIntro()

	for {
		showOptions()
		option := readOption()

		switch option {
		case 1:
			fmt.Println("Starting monitoring...")
			monitoring()
		case 2:
			fmt.Println("Showing logs...")
			showLogs()
		case 0:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid option")
			os.Exit(-1)
		}
	}
}
