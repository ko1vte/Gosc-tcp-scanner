package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"sort"
)

var addr string
var portnum int
var goroutine int

func main() {
	cmd()
}
func cmd() {
	//ports := make(chan int, 100)
	flag.IntVar(&portnum, "p", 1024, "port number Default:1024")
	flag.IntVar(&goroutine, "g", 100, "goroutine Default:100")
	flag.Parse()

	for i := 1; i < len(os.Args); i++ {
		addr = os.Args[i]
	}

	if ipv4Addr(addr) {
		fmt.Printf("port number:%v\naddress:%v\n[+] Port Scanning\n", portnum, addr)
		scanner(goroutine, portnum, addr)

	} else {
		fmt.Println("The IP you entered is not IPV4")
	}

}

func ipv4Addr(ip string) bool {
	address := net.ParseIP(ip)
	if address == nil {
		return false
	} else {
		return true
	}
}

func worker(ports chan int, address string, results chan int) {
	for p := range ports {
		addr := fmt.Sprintf("%v:%d", address, p)
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func scanner(goroutine int, portnum int, addr string) {
	ports := make(chan int, goroutine)
	results := make(chan int)
	var openports []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, addr, results)
	}
	go func() {
		for i := 1; i < portnum; i++ {
			ports <- i
		}
	}()

	for i := 1; i < portnum; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}
	close(ports)
	close(results)

	sort.Ints(openports)

	for _, openports := range openports {
		fmt.Printf("[+] Port %d is open\n", openports)
	}
	fmt.Println("[-] Port Scanning Completed")
}
