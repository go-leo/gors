package parser

import (
	"fmt"
	"testing"
)

func TestFindWildcard(t *testing.T) {
	path := "/v1/:version/user/:id/book"

	wildcard, i, valid := FindWildcard(path)
	fmt.Println(wildcard, i, valid)
	path = path[i+len(wildcard):]
	wildcard, i, valid = FindWildcard(path)
	fmt.Println(wildcard, i, valid)
	path = path[i+len(wildcard):]
	wildcard, i, valid = FindWildcard(path)
	fmt.Println(wildcard, i, valid)

}

func TestFindWildcards(t *testing.T) {
	path := "/v1/:version/user/:id/book"
	wildcards, err := FindWildcards(path)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(wildcards)
}
