package main

import (
	"fmt"

	chuck "github.com/pjsoftware/go-api-chuck"
)

type Joke struct {
	Value string
}

func main() {
	cn := chuck.New()
	fmt.Println("Chuck [Auth0: No Authorisation]")

	cj := cn.RandomJoke()
	fmt.Printf("\n%s\n\n", cj.Value)
	fmt.Printf("- ID: \"%s\"\n", cj.ID)
	if len(cj.Categories) > 0 {
		fmt.Printf("- Categories: %v\n", cj.Categories)
	}

	cl := cn.CategoryList()
	fmt.Printf("\nCategories: %v\n", cl)
}
