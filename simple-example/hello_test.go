package hello

import (
	"testing"
)

func TestSayHello(t *testing.T) {
	subtests := []struct {
		items  []string
		result string
	}{
		{
			result: "Hello, World!",
		},
		{
			items:  []string{"Rubem"},
			result: "Hello, Rubem!",
		},
		{
			items:  []string{"Rubem", "Eduarda"},
			result: "Hello, Rubem, Eduarda!",
		},
	}

	for _, st := range subtests {
		if s := Say(st.items); s != st.result {
			t.Errorf(("wanted %s (%v), got %s"), st.items, st.result, s)
		}
	}
}
