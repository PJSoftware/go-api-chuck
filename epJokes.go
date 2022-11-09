package chuck

import (
	"encoding/json"
)

// RandomJoke returns a random joke from the Chuck Norris API
func (c *ChuckAPI) RandomJoke() *Joke {
	res, err := c.api.Get("/random")
	chuckHatesErrors(err)

	joke := &Joke{}
	json.Unmarshal(res.Body, joke)

	return joke
}
