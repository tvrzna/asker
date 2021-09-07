package main

import (
	"fmt"
	"os"

	"github.com/tvrzna/go-utils/args"
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

	if args.ContainsArg(os.Args, "-h", "--help") {
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
		os.Exit(0)
	}

	if args.ContainsArg(os.Args, "-v", "--version") {
		fmt.Println(getVersion())
		os.Exit(0)
	}

	args.ParseArgs(os.Args, func(arg, nextArg string) {
		switch arg {
		case "-t", "--title":
			conf.title = nextArg
		case "-m", "--message":
			conf.message = nextArg
		case "-c", "--confirm":
			conf.msgType = confirm
		case "-w", "--warning":
			conf.msgType = warning
		case "-i", "--info":
			conf.msgType = info
		case "-p", "--pasword":
			conf.msgType = password
		}
	})

	return conf
}
