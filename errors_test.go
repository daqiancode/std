package std_test

import (
	"fmt"
	"testing"

	"github.com/daqiancode/std"
)

func TestError(t *testing.T) {
	ok := std.OK
	fmt.Println(ok.ToResult(nil))
}
