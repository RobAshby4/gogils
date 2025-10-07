package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

var mediaWikiApiUrl string = "https://ffxiv.consolegameswiki.com/mediawiki/api.php?action=parse&prop=wikitext&format=json&page="

func getItemPageResp(item Item) ([]byte, error) {
	var contents []byte
	queryUrl := mediaWikiApiUrl + item.wikiName

	resp, err := http.Get(queryUrl)
	if err != nil {
		return contents, fmt.Errorf("failed to fetch from %s", queryUrl)
	}
	if resp.StatusCode != 200 {
		return contents, fmt.Errorf("failed to fetch from %s, status code %d", queryUrl, resp.StatusCode)
	}
	defer resp.Body.Close()

	contents, err = io.ReadAll(resp.Body)
	if err != nil {
		return contents, fmt.Errorf("failed to read response from %s", queryUrl)
	}

	return contents, nil
}

type WikiText struct {
	Text string `json:"*"`
}

type Parse struct {
	Title    string   `json:"title"`
	PageID   int      `json:"pageid"`
	WikiText WikiText `json:"wikitext"`
}

type MediaWikiResponse struct {
	Parse Parse `json:"parse"`
}

func extractArticleFromBody(apiResponse []byte) (string, error) {
	var article string
	var result MediaWikiResponse
	err := json.Unmarshal(apiResponse, &result)
	if err != nil {
		return article, err
	}
	article = result.Parse.WikiText.Text
	return article, err
}

func articleContainsRecipe(lines []string) bool {
	for _, line := range lines {
		if strings.Contains(line, "{{Recipe") {
			return true
		}
	}
	return false
}

func extractAllIngredients(lines []string) map[string]int {
	reFindIngredient := regexp.MustCompile(`^\|\s*ingredient\d+\s*=`)
	reParseIngredient := regexp.MustCompile(`=\s*(\d+)\s+(.+)$`)
	var recipe map[string]int = make(map[string]int)
	for _, line := range lines {
		if reFindIngredient.MatchString(line) {
			ingredient, quantity := processIngredientLine(line, reParseIngredient)
			recipe[ingredient] = quantity
		}
	}
	return recipe
}

func processIngredientLine(line string, re *regexp.Regexp) (string, int) { // returns ingredient and quantity
	matches := re.FindStringSubmatch(line)
	var quantity int
	var ingredient string
	if len(matches) == 3 {
		quantity, _ = strconv.Atoi(matches[1])
		ingredient = matches[2]
	}
	return ingredient, quantity
}

func GetRecipeFromWiki(item *Item) (map[string]int, error) {
	recipe := make(map[string]int)
	apiResponse, err := getItemPageResp(*item)
	if err != nil {
		return recipe, err
	}
	article, err := extractArticleFromBody(apiResponse)
	if err != nil {
		return recipe, err
	}
	articleLines := strings.Split(article, "\n")
	if articleContainsRecipe(articleLines) {
		// search for ingredients and add them to the recipe
		extractedRecipe := extractAllIngredients(articleLines)
		for itemName, quantity := range extractedRecipe {
			if len(itemName) > 0 {
				recipe[itemName] = quantity
			}
		}
	} else {
		err = fmt.Errorf("no recipe found on page")
	}
	// use MediaWiki API to fetch recipe, if not existant, return err
	fmt.Println(recipe)
	return recipe, err
}
