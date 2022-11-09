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
	fmt.Println(cn.Ident())

	cj := cn.RandomJoke()
	fmt.Printf("\n%s\n\n", cj.Value)
	fmt.Printf("- ID: \"%s\"\n", cj.ID)
	if len(cj.Categories) > 0 {
		fmt.Printf("- Categories: %v\n", cj.Categories)
	}

	cl := cn.CategoryList()
	fmt.Printf("\nCategories: %v\n", cl)

	cc := cn.RandomByCategory("food")
	fmt.Printf("\n%s\n\n", cc.Value)
	fmt.Printf("- ID: \"%s\"\n", cc.ID)
	if len(cc.Categories) > 0 {
		fmt.Printf("- Categories (incl FOOD?): %v\n", cc.Categories)
	}

	cs := cn.SearchFor("food")
	fmt.Printf("\nFound %d jokes containing 'food':\n", cs.Total)

	for idx, joke := range cs.Result {
		fmt.Printf("%02d: %s\n", idx+1, joke.Value)
	} 
}
