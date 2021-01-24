package main

import (
	"flag"
	"fmt"
	"os"
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
	conf := &config{}

	var showConfirm bool
	var showWarning bool
	var showInfo bool
	var showPassword bool
	var title string
	var message string

	var printHelp bool
	var printVersion bool

	flag.BoolVar(&printHelp, "h", false, "print this help")

	flag.BoolVar(&printVersion, "v", false, "print version")

	flag.StringVar(&title, "t", "asker", "")

	flag.StringVar(&message, "m", "", "")

	flag.BoolVar(&showConfirm, "c", false, "")

	flag.BoolVar(&showWarning, "w", false, "")

	flag.BoolVar(&showInfo, "i", false, "")
	flag.BoolVar(&showInfo, "-info", false, "")

	flag.BoolVar(&showPassword, "p", false, "")

	flag.Parse()

	if printHelp {
		fmt.Println("Usage: asker [options]")
		fmt.Println("Options:")
		fmt.Printf("  -h\t\tprint this help\n")
		fmt.Printf("  -v\t\tprint version\n")
		fmt.Printf("  -t TEXT\t\tsets the title\n")
		fmt.Printf("  -m TEXT\t\tsets the message\n")
		fmt.Printf("  -c\t\tshows confirm dialog \n")
		fmt.Printf("  -i\t\tshows info dialog \n")
		fmt.Printf("  -w,\t\tshows warning dialog \n")
		fmt.Printf("  -p\t\tshows password dialog \n")

		os.Exit(0)
	}

	if printVersion {
		fmt.Println(getVersion())
		os.Exit(0)
	}

	if showPassword {
		conf.msgType = password
	} else if showWarning {
		conf.msgType = warning
	} else if showInfo {
		conf.msgType = info
	} else {
		conf.msgType = confirm
	}

	conf.title = title
	conf.message = message

	return conf
}
