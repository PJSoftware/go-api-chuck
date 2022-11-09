package chuck

import (
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

	// initialise all endpoints
	chuck.epRandom = chuck.api.NewEndpoint("/random")
	chuck.epCategories = chuck.api.NewEndpoint("/categories")
	chuck.epSearch = chuck.api.NewEndpoint("/search")
	return chuck
}

func chuckHatesErrors(err error) {
	if err != nil {
		log.Fatalf("Error? Chuck Norris does not make errors. Therefore this must be your fault: %v", err)
	}
}