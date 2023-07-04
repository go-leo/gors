package parser_test

import (
	"fmt"
	"testing"
)

func TestExtractValue(t *testing.T) {
	value, b := parser.ExtractValue("@Path", parser.Path)
	fmt.Println(value, b)

	value, b = parser.ExtractValue("@Path11", parser.Path)
	fmt.Println(value, b)
	value, b = parser.ExtractValue("@Path(", parser.Path)
	fmt.Println(value, b)
	value, b = parser.ExtractValue("@Path( ", parser.Path)
	fmt.Println(value, b)
	value, b = parser.ExtractValue("@Path()", parser.Path)
	fmt.Println(value, b)
	value, b = parser.ExtractValue("@Path( )", parser.Path)
	fmt.Println(value, b)

	value, b = parser.ExtractValue("@Path( ) ", parser.Path)
	fmt.Println(value, b)

	value, b = parser.ExtractValue("@Path(dddd ) ", parser.Path)
	fmt.Println(value, b)
}
