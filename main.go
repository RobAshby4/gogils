package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner bufio.Scanner

func main() {
	initGlobals()
	runloop()
}

func initGlobals() {
	scanner = *bufio.NewScanner(os.Stdin)
}

// TODO: breakout into smaller funcs
func queryItems(itemLog ItemLog) {
	queryQueue := make([]*Item, 0)
	fmt.Print("query > ")
	for scanner.Scan() {
		// reset output buffer for new ingredient
		outputBuffer := ""
		completedQueryQueue := make([]*Item, 1)
		line := scanner.Text()
		item, err := itemLog.GetItemByName(line)
		if err != nil {
			BPrintln(&outputBuffer, "item \""+line+"\" not found")
			if len(queryQueue) > 0 {
				queryQueue = queryQueue[1:]
			}
			continue
		}
		queryQueue = append(queryQueue, item)
		fmt.Println(completedQueryQueue)
		for len(queryQueue) > 0 {
			item = queryQueue[0]
			fmt.Print(queryQueue[0])
			fmt.Println(" running in queue")
			// item exhists
			if item.IsCraftable() {
				// if craftable, print ingredients and add them to the queue if not already queued
				BPrintln(&outputBuffer, item.name+" is craftable. recipe:")
				for ingredientName, ingredientQuantity := range *item.recipe {
					BPrintln(&outputBuffer, strconv.Itoa(ingredientQuantity)+" | "+ingredientName)
					ingredientItem, err := itemLog.GetItemByName(ingredientName)
					if err != nil {
						BPrintln(&outputBuffer, "ingredient \""+line+"\" not found")
						continue
					}
					if ingredientItem.IsCraftable() {
						if !SliceContainsPtr(completedQueryQueue, ingredientItem) {
							queryQueue = append(queryQueue, ingredientItem)
							completedQueryQueue = append(completedQueryQueue, ingredientItem)
							fmt.Println(completedQueryQueue)
							fmt.Println(ingredientItem)
							fmt.Println("added " + ingredientItem.name)
						}
					}
				}

			} else {
				BPrintln(&outputBuffer, item.name+" is NOT craftable: No recipe found.")
			}
			if len(queryQueue) > 0 {
				queryQueue = queryQueue[1:]
			}
		}
		BPrint(&outputBuffer, "query > ")
		// draw output from previous request
		print(outputBuffer)
	}
}

func runloop() {
	itemLog := GetItemLog()
	queryItems(*itemLog)
}
