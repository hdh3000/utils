package markov

import (
	"fmt"
	"testing"
)

func TestMakeID(t *testing.T) {
	key, sum := makeKey([]string{"foo", "bar", "boasdfasdfsdfom"})
	fmt.Printf("%d, %d\n", len(key), sum)
}
