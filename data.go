package chuck

type JokeCategories []string

type Joke struct {
	Value      string
	ID         string
	Categories JokeCategories
	URL        string
	IconURL    string `json:"icon_url"`
	Created    string `json:"created_at"`
	Updated    string `json:"updated_at"`
}

type Jokes struct {
	Total  int
	Result []Joke
}
