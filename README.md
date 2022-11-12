# go-api-chuck

The [Chuck Norris API](https://api.chucknorris.io) is a simple web API with
several endpoints, which returns Chuck Norris jokes. It has no authentication
requirements, and hence is a good starting point for the development of the
[go-api](https://github.com/PJSoftware/go-api) package.

This `go-api-chuck` package provides an interface for the `Chuck Norris API`,
using `go-api` to communicate with the API.

## Chuck Norris API

### EndPoints

- [x] `/random` returns a random joke.
- [x] `/categories` returns a list of available categories.
- [x] `/random?category={category}` returns a random joke from the named category.
- [x] `/search?query={query}` returns a list of jokes matching the query string.

## cmd/chuck

The `chuck` command uses `go-api-chuck` to query the `Chuck Norris API`. You
can run it with:

```sh
go run ./cmd/chuck
```

Alternatively, if you have [task](https://taskfile.dev) installed, you can use:

```sh
task chuck
```

### Command Line Flags

- `bin/chuck.exe` returns a random Chuck Norris joke.
- `bin/chuck.exe -list` returns a list of available categories.
- `bin/chuck.exe -cat <cat>` returns a random joke from the specified category.
- `bin/chuck.exe -find <str>` lists all jokes containing the specified string.
