package main

import (
	"fmt"
	"os"
	"strings"
)

type enMsgType int

const (
	confirm enMsgType = iota
	info
	warning
	password
)

type config struct {
	msgType enMsgType
	title   string
	message string
}

func loadConfig() *config {
	conf := &config{msgType: warning, title: "asker", message: "Please see help how to use this application."}

	if contains(os.Args, "-h") || contains(os.Args, "--help") {
		printHelp()
		os.Exit(0)
	}
	if contains(os.Args, "-v") || contains(os.Args, "--version") {
		fmt.Println(getVersion())
		os.Exit(0)
	}

	for i, arg := range os.Args {
		switch arg {
		case "-t", "--title":
			nextArg(os.Args, i, func(val string) {
				conf.title = strings.TrimSpace(val)
			})
		case "-m", "--message":
			nextArg(os.Args, i, func(val string) {
				conf.message = strings.TrimSpace(val)
			})
		case "-c", "--confirm":
			conf.msgType = confirm
		case "-w", "--warning":
			conf.msgType = warning
		case "-i", "--info":
			conf.msgType = info
		case "-p", "--pasword":
			conf.msgType = password
		}
	}

	return conf
}

// Checks, if array contains value
func contains(array []string, value string) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

// Gets next argument, if available
func nextArg(args []string, i int, callback func(value string)) {
	if callback != nil && len(args) > i+1 {
		val := strings.TrimSpace(args[i+1])
		if !strings.HasPrefix(val, "-") {
			callback(args[i+1])
		}
	}
}

func printHelp() {
	fmt.Println("Usage: asker [options]")
	fmt.Println("Options:")
	fmt.Printf("  -h, --help\t\t\tprint this help\n")
	fmt.Printf("  -v, --version\t\t\tprint version\n")
	fmt.Printf("  -t TEXT, --title TEXT\t\tsets the title\n")
	fmt.Printf("  -m TEXT, --message TEXT\tsets the message\n")
	fmt.Printf("  -c, --confirm\t\t\tshows confirm dialog \n")
	fmt.Printf("  -i, --info\t\t\tshows info dialog \n")
	fmt.Printf("  -w, --warning\t\t\tshows warning dialog \n")
	fmt.Printf("  -p, --password\t\tshows password dialog \n")
}
