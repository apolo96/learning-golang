# Golang - Modules Manager

## Structure 

Repository > Module > Package



A repository is a place in a version control system where the source code for a project is stored. 

A module is a bundle of Go source code that's distributed and versioned as a single unit. Modules consist of one or more packages, which are directories of source code. 

An Packages give a module organization and structure.

Module PATH:
```
REPOSITORY_NAME/OWNER/MODULE_NAME
```
Example:
```
github.com/apolo96/regulax
```

## How using a thridparty Module in golang?

In your code import the module, example:
```
import "github.com/learning-go-book-2e/simpletax"
```

Then, download the module:
```
go get github.com/learning-go-book-2e/simpletax
```
By default, go tool downloads the last version of module.

## Manage Module Versions

List the module versions
```
go list -m -versions github.com/learning-go-book-2e/simpletax
```

```
go get ./..
```

```
go get -u=patch $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}@{{.Version}}{{end}}' -m all)
```

```
go clean -modcache
```

```
go mod tidy
```