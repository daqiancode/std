package std_test

import (
	"fmt"
	"testing"

	"github.com/daqiancode/std"
)

func TestPage(t *testing.T) {
	s := "a asc,b,c"
	orderBys := std.ParseOrderBys(s)
	r := orderBys.Pick("a", "b").AddPrefix(map[string]string{"a": "b"})
	fmt.Println(r)
}
