// You want to log messages in different places:
// Console
// File
// Remote API
// Every logger must have a Log(message string) method.
// ✅ Design the system using interfaces so the logging mechanism can be swapped easily.
package main

import "fmt"

type logging_interface interface {
	Log(message string)
}

type console_log struct {
	location string
}

func (c console_log) Log(message string) {
	fmt.Println("log from: ", c.location, " message: ", message)
}

type file_log struct {
	location string
}

func (f file_log) Log(message string) {
	fmt.Println("log from: ", f.location, " message: ", message)
}

func binder(l logging_interface, message string) {
	l.Log(message)
}

func main() {

	var l logging_interface
	l = console_log{
		location: "/var/log/",
	}
	binder(l, " this is  log of console")
	l = file_log{
		location: "/var/file",
	}
	binder(l, " this is  log of file")

}
