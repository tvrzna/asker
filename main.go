package main

import (
	"fmt"
	"os"
)

const version = "0.2.0"

func main() {
	result, text := handleDialog(loadConfig())
	if result == yes {
		fmt.Println(text)
		os.Exit(0)
		return
	} else {
		os.Exit(1)
		return
	}

}

func getVersion() string {
	return fmt.Sprintf("asker %s\nhttps://github.com/tvrzna/asker\n\nReleased under the MIT License.", version)
}
