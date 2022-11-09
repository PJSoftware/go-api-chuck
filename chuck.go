package chuck

import (
	"log"

	goapi "github.com/pjsoftware/go-api"
)

type ChuckAPI struct {
	api *goapi.APIData
}

func New() *ChuckAPI {
	chuck := &ChuckAPI{}
	chuck.api = goapi.New("http://api.chucknorris.io/jokes")
	return chuck
}

func chuckHatesErrors(err error) {
	if err != nil {
		log.Fatalf("Error? Chuck Norris does not make errors. Therefore this must be your fault: %v", err)
	}
}