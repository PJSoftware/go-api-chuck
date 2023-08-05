package chuck

import (
	"encoding/json"
)

// RandomJoke returns a random joke from the Chuck Norris API
func (c *ChuckAPI) RandomJoke() *Joke {
	res, err := c.epRandom.NewRequest().GET()
	chuckHatesErrors(res.Status, err)

	joke := &Joke{}
	json.Unmarshal([]byte(res.Body), joke)

	return joke
}
