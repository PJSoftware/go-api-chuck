package chuck

import (
	"encoding/json"
)

// RandomJoke returns a random joke from the Chuck Norris API
func (c *ChuckAPI) RandomByCategory(cat string) *Joke {
	if !c.ExistsCategory(cat) {
		return nil
	}

	req := c.epRandom.NewRequest()
	req.AddQuery("category", cat)
	res, err := req.GET()
	chuckHatesErrors(err)

	joke := &Joke{}
	json.Unmarshal(res.Body, joke)

	return joke
}
