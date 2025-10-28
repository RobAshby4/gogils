package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type ItemLog struct {
	// TODO: Refactor into a map
	items    []Item
	numItems int
}

var itemLog *ItemLog

func GetItemLog() *ItemLog {
	if itemLog == nil {
		InitItemLog()
	}
	return itemLog
}

func InitItemLog() {
	// read items file into byte array
	itemFile, err := os.Open("./items.json")
	if err != nil {
		panic(err)
	}
	defer itemFile.Close()

	itemBytes, err := io.ReadAll(itemFile)
	if err != nil {
		panic(err)
	}

	// create variable to match json input
	var jsonImport map[string]map[string]string
	// load byte array into variable
	json.Unmarshal(itemBytes, &jsonImport)

	// extract data from json, remove blank items
	var itemsSlice []Item
	numItems := 0
	for idString, itemData := range jsonImport {
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}
		name := itemData["en"]
		if len(name) > 0 {
			itemsSlice = append(itemsSlice, NewItem(id, name))
			numItems++
		}
	}

	// create and assign ItemLog
	itemLog = &ItemLog{itemsSlice, numItems}
}

func (itemLog *ItemLog) GetItemByID(id int) (*Item, error) {
	for _, item := range itemLog.items {
		if item.id == id {
			return &item, nil
		}
	}

	// no matches found :(
	return nil, fmt.Errorf("id %d not found", id)
}

func (itemLog *ItemLog) GetItemByName(query string) (*Item, error) {
	for _, item := range itemLog.items {
		// if substring is present, add to matches
		if strings.EqualFold(item.name, query) {
			return &item, nil
		}
	}

	// no matches found :(
	return nil, fmt.Errorf("%s not found", query)
}

/* deprecated
func (itemLog *ItemLog) getRecipeItems(recipe map[string]int) []*Item {
	items := make([]*Item, 0)
	for itemName := range recipe {
		item, err := itemLog.GetItemByName(itemName)
		if err != nil {
			fmt.Println("Item not found " + itemName)
			continue
		}
		items = append(items, item)
	}
	return items
}
*/
