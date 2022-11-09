package chuck

import (
	"encoding/json"
)

// CategoryList returns a list of valid categories from the Chuck Norris API
func (c *ChuckAPI) CategoryList() *JokeCategories {
	res, err := c.epCategories.NewRequest().GET()
	chuckHatesErrors(err)

	catList := &JokeCategories{}
	json.Unmarshal(res.Body, catList)

	return catList
}
