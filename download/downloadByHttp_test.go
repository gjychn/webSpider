package download

import (
	"fmt"
	"testing"
)

type aa struct {
	a int
	b string
}

func TestA(t *testing.T) {
	a := aa{
		12,
		"sssss",
	}
	s := fmt.Sprintf("I am %T", a)
	fmt.Println(s)
}
