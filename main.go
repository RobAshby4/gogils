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

func runloop() {
	itemLog := GetItemLog()
	fmt.Print("query > ")
	for scanner.Scan() {
		line := scanner.Text()
		item, _ := itemLog.GetItemByName(line)
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
