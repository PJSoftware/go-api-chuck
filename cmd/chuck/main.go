package main

import (
	"flag"
	"fmt"

	chuck "github.com/pjsoftware/go-api-chuck"
)

type Joke struct {
	Value string
}

var cnAPI *chuck.ChuckAPI

func main() {
	flgCatList := flag.Bool("list", false, "return a list of categories")
	flgSearch := flag.String("find", "", "specify a search string")
	flgInCat := flag.String("cat", "", "specify a category to be used")
	flag.Parse()

	cnAPI = chuck.New()
	fmt.Println(cnAPI.Ident())

	if *flgCatList {
		catList()
		return
	}

	if *flgSearch != "" {
		searchFor(*flgSearch)
		return
	}

	if *flgInCat != "" {
		searchCategory(*flgInCat)
		return
	}

	randomJoke()
}

func randomJoke() {
	fmt.Printf("Random Joke:\n")
	cj := cnAPI.RandomJoke()
	fmt.Printf("\n%s\n", cj.Value)
	fmt.Println("")
}

func catList() {
	fmt.Printf("Categories:\n")
	cl := cnAPI.CategoryList()
	for _, cat := range *cl {
		fmt.Printf("- %s\n", cat)
	}
	fmt.Println("")
}

func searchFor(str string) {
	fmt.Printf("Search for '%s':\n", str)
	cs := cnAPI.SearchFor(str)
	fmt.Printf("\nFound %d jokes containing '%s':\n", cs.Total, str)

	for idx, joke := range cs.Result {
		fmt.Printf("%02d: %s\n", idx+1, joke.Value)
	}
	fmt.Println("")
}

func searchCategory(cat string) {
	fmt.Printf("Random joke in '%s' category:\n", cat)
	cc := cnAPI.RandomByCategory(cat)
	if cc == nil {
		fmt.Printf("* Category '%s' does not exist\n", cat)
		fmt.Printf("* See 'chuck -list' for more!\n")
		return
	}

	fmt.Printf("\n%s\n", cc.Value)
	fmt.Println("")
}
