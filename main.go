package main

import (
	"bufio"
	"fmt"
	"os"
)

var scanner bufio.Scanner

func main() {
	initGlobals()
	runloop()
}

func initGlobals() {
	scanner = *bufio.NewScanner(os.Stdin)
}

func queryItems(itemLog ItemLog) {
	// queryQueue := make([]Item, 1) will use a queue once supporting functions are added to items
	fmt.Print("query > ")
	for scanner.Scan() {
		line := scanner.Text()
		item, err := itemLog.GetItemByName(line)
		if err != nil {
			fmt.Println("item \"" + line + "\" not found")
			fmt.Print("\nquery > ")
			continue
		}
		if item.IsCraftable() {
			recipeItems := itemLog.getRecipeItems(*item.recipe)
			for _, recipeItem := range recipeItems {
				fmt.Println(*recipeItem)
			}
			fmt.Println(*item)
		}
		fmt.Print("query > ")
	}
}

func runloop() {
	itemLog := GetItemLog()
	queryItems(*itemLog)
}
