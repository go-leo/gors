package annotation_test

import (
	"fmt"
	"github.com/go-leo/gors/internal/pkg/annotation"
	"testing"
)

func TestExtractValue(t *testing.T) {
	value, b := annotation.ExtractValue("@Path", annotation.Path)
	fmt.Println(value, b)

	value, b = annotation.ExtractValue("@Path11", annotation.Path)
	fmt.Println(value, b)
	value, b = annotation.ExtractValue("@Path(", annotation.Path)
	fmt.Println(value, b)
	value, b = annotation.ExtractValue("@Path( ", annotation.Path)
	fmt.Println(value, b)
	value, b = annotation.ExtractValue("@Path()", annotation.Path)
	fmt.Println(value, b)
	value, b = annotation.ExtractValue("@Path( )", annotation.Path)
	fmt.Println(value, b)

	value, b = annotation.ExtractValue("@Path( ) ", annotation.Path)
	fmt.Println(value, b)

	value, b = annotation.ExtractValue("@Path(dddd ) ", annotation.Path)
	fmt.Println(value, b)
}
