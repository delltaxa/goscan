package main

import "fmt"

type _icons struct {
	_event string
	_info string
	_warning string
	_error string
}

var log_icon _icons = _icons {
	_event: "\x1b[34m[*]\x1b[39m",
	_info: "\x1b[32m[+]\x1b[39m",
	_warning: "\x1b[33m[!]\x1b[39m",
	_error: "\x1b[31m[-]\x1b[39m",
}

func log_msg(ico string, msg string) {
	fmt.Printf(ico + " " + msg)
}