package main

var logo string = Fore["BLUE"] + `
                           `+Fore["CYAN"]+`   ___  `+Fore["BLUE"]+`
 ___ ___ ___ ___ ___ ___   `+Fore["CYAN"]+`  /   \ `+Fore["BLUE"]+`
| . | . |_ -|  _| .'|   |  `+Fore["CYAN"]+` | o o |`+Fore["BLUE"]+`
|_  |___|___|___|__,|_|_|  `+Fore["CYAN"]+` |  _  |`+Fore["BLUE"]+`
|___|    `+Fore["YELLOW"]+`v1.0.0.0`+Fore["CYAN"]+`           (_) (_)

`

type _cfg struct {
	addr string
	ports []int
	protocol string
}

var cfg _cfg = _cfg {
	addr: "192.168.178.175",
	protocol: "tcp",
	ports: []int {80, 443, 1330, 123, 1999},
}
