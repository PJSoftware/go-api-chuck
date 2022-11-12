package chuck

import (
	"encoding/json"
	"strings"
)

// CategoryList returns a list of valid categories from the Chuck Norris API
func (c *ChuckAPI) CategoryList() *JokeCategories {
	res, err := c.epCategories.NewRequest().GET()
	chuckHatesErrors(err)

	catList := &JokeCategories{}
	json.Unmarshal(res.Body, catList)

	return catList
}

// ExistsCategory returns true if the category exists; otherwise it returns false
func (c *ChuckAPI) ExistsCategory(catInQuestion string) bool {
	catList := c.CategoryList()
	for _, cat := range *catList {
		if strings.EqualFold(cat, catInQuestion) {
			return true
		}
	}

	return false
}
