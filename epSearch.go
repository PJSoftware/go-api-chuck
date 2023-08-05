package chuck

import (
	"encoding/json"
)

// SearchFor returns a list of all jokes containing the search string
func (c *ChuckAPI) SearchFor(str string) *Jokes {
	req := c.epSearch.NewRequest()
	req.AddQuery("query", str)
	res, err := req.GET()
	chuckHatesErrors(err)

	jokes := &Jokes{}
	json.Unmarshal([]byte(res.Body), jokes)

	return jokes
}
