package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var discovered_ports []int

func scanPort(protocol, hostname string, port int) bool {
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 5*time.Second)

	if err != nil {
		return false
	}

	defer conn.Close()
	
	return true
}

func worker(wg *sync.WaitGroup, id int, check bool) {
	defer wg.Done()

	open := scanPort(cfg.protocol, cfg.addr, id)
	_ = open

	if open {
		if check == false {
			log_msg(log_icon._info, "Discovered open port " + Fore["GREEN"] + strconv.Itoa(id) + Fore["RESET"] + "\n")
			discovered_ports = append(discovered_ports, id)
		} else {
			log_msg(log_icon._info, "Open Port: " + Fore["GREEN"] + strconv.Itoa(id) + Fore["RESET"] + "\n")
		}
	}

}

func portRange(min int, max int) []int {
	if min > max {
		return []int {}
	}
	
	var r []int
	for i:=min;i<max;i++ {
		r = append(r, i)
	}

	return r
}

func parse_args() {
	var args []string = os.Args

	if len(args) > 1 {
		var argin string = strings.ToLower(args[1])

		if strings.HasPrefix(argin, "tcp://") {
			cfg.protocol = "tcp"
		} else if strings.HasPrefix(argin, "udp://") {
			cfg.protocol = "udp"
		} else {
			log_msg(log_icon._error, "Only tcp/udp protocols are supported")
			os.Exit(0)
		}

		argin = strings.ReplaceAll(argin, "tcp://", "")
		argin = strings.ReplaceAll(argin, "udp://", "")

		if strings.Contains(argin, ":") {
			cfg.addr = strings.Split(argin, ":")[0]
		} else {
			log_msg(log_icon._error, "Missing port")
			os.Exit(0)
		}

		argin = strings.ReplaceAll(argin, cfg.addr+":", "")

		if strings.Contains(argin, "-") {
			min, errn := strconv.Atoi(strings.Split(argin, "-")[0])
			max, errx := strconv.Atoi(strings.Split(argin, "-")[1])
			
			if errn != nil || errx != nil {
				log_msg(log_icon._error, "Port range is invalid")
				os.Exit(0)	
			}

			cfg.ports = portRange(min, max)
		} else {
			log_msg(log_icon._error, "Port range is invalid")
			os.Exit(0)
		}
		

	} else {
		log_msg(log_icon._error, "No args received")
		os.Exit(0)
	}
}

func main() {
	fmt.Printf(logo)

	parse_args()

	log_msg(log_icon._event, "Starting port scan for ("+Fore["BLUE"]+cfg.addr+Fore["RESET"]+")\n")

	var wg sync.WaitGroup
    var maxWorkers = len(cfg.ports)

	log_msg(log_icon._event, "Adding Threads\n\n")
    for i := 0; i < maxWorkers; i++ {
		// add thread
        wg.Add(1)

        go worker(&wg, cfg.ports[i], false)
    }	

    wg.Wait()

	if len(discovered_ports) == 0 {
		log_msg(log_icon._error, "No ports discovered")
	} else {
		fmt.Println()
		log_msg(log_icon._event, "Checking discovered Ports\n")
		var wg sync.WaitGroup
		var maxWorkers = len(discovered_ports)

		log_msg(log_icon._event, "Adding Threads\n")
		for i := 0; i < maxWorkers; i++ {
			wg.Add(1)

			go worker(&wg, discovered_ports[i], true)
		}	

		fmt.Printf("\n")

		wg.Wait()
	}

}
