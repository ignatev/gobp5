package main

import(
	"os"
	"fmt"
	"net"
	"log"
//	"time"
	"bufio"
	"strings"
)

var marks = map[bool]string{true: "✔", false: "✖"}

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		domain := s.Text()
		fmt.Print(domain, " ")
		exists, err := exists(domain)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(marks[!exists])
		//time.Sleep(1 * time.Second)
	}
}

func exists(domain string) (bool, error) {
	const whoisServer string = "com.whois-servers.net"
	conn, err := net.Dial("tcp", whoisServer+":43")
	if err != nil {
		return false, err
	}
	defer conn.Close()
	conn.Write([]byte(domain + "rn"))
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		if strings.Contains(strings.ToLower(scanner.Text()), "no match") {
			return false, nil
		}
	}
	return true, nil
}