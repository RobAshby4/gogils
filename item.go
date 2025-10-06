package main

import (
	"fmt"
	"strings"
)

type Item struct {
	id       int
	name     string
	wikiName string
	// these are pointers because they may be nil
	craftable    *bool
	recipe       *map[string]int // name, num required
	purchaseable bool
	price        int
}

func NewItem(id int, name string) Item {
	var newItem Item
	newItem.id = id
	newItem.name = name
	newItem.wikiName = strings.ReplaceAll(name, " ", "_")
	return newItem
}

func (item *Item) IsCraftable() bool {
	if item.craftable == nil {
		val := true
		recipe, err := GetRecipeFromWiki(*item)
		if err != nil {
			item.craftable = &val
			fmt.Println(err)
		} else {
			item.craftable = &val
			item.recipe = &recipe
		}
	}
	return *item.craftable
}
