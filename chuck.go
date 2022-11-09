package chuck

import (
	"fmt"
	"log"

	goapi "github.com/pjsoftware/go-api"
)

type ChuckAPI struct {
	api 					*goapi.APIData
	epRandom 			*goapi.Endpoint
	epCategories 	*goapi.Endpoint
	epSearch 	    *goapi.Endpoint
}

func New() *ChuckAPI {
	chuck := &ChuckAPI{}
	chuck.api = goapi.New("http://api.chucknorris.io/jokes")
	chuck.api.SetName(fmt.Sprintf("Chuck API v%s", pkgVersion))

	// initialise all endpoints
	chuck.epRandom = chuck.api.NewEndpoint("/random")
	chuck.epCategories = chuck.api.NewEndpoint("/categories")
	chuck.epSearch = chuck.api.NewEndpoint("/search")
	return chuck
}

func (c *ChuckAPI) Ident() string {
	return c.api.Ident()
}

func chuckHatesErrors(err error) {
	if err != nil {
		log.Fatalf("Error? Chuck Norris does not make errors. Therefore this must be your fault: %v", err)
	}
}