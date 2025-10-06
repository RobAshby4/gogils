package main

import (
	"bufio"
	"fmt"
	"os"
)

var itemLog ItemLog
var scanner bufio.Scanner

func main() {
	initGlobals()
	runloop()
}

func initGlobals() {
	itemLog = NewItemlog()
	scanner = *bufio.NewScanner(os.Stdin)
}

func runloop() {
	for scanner.Scan() {
		line := scanner.Text()
		item, _ := itemLog.GetItemByName(line)
		fmt.Println(item[0].IsCraftable())
	}
}
