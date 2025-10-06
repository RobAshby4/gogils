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

func NewItemlog() ItemLog {
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

	// create and return ItemLog
	itemLog := ItemLog{itemsSlice, numItems}
	return itemLog
}

func (itemLog *ItemLog) GetItemByID(id int) (Item, error) {
	for _, item := range itemLog.items {
		if item.id == id {
			return item, nil
		}
	}

	// no matches found :(
	return Item{}, fmt.Errorf("id %d not found", id)
}

func (itemLog *ItemLog) GetItemByName(query string) ([]Item, error) {
	var matchedItems []Item
	for _, item := range itemLog.items {
		// if substring is present, add to matches
		if strings.Contains(strings.ToLower(item.name), strings.ToLower(query)) {
			matchedItems = append(matchedItems, item)
		}
	}

	if len(matchedItems) > 0 {
		return matchedItems, nil
	}

	// no matches found :(
	return matchedItems, fmt.Errorf("%s not found", query)
}
