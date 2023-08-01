package hello

import "testing"

func TestSayHello(t *testing.T) {
	subtests := []struct {
		items  []string
		result string
	}{
		{
			result: "Hello, world!",
		},
		{
			result: "Hello, mark!",
			items:  []string{"mark"},
		},
		{
			result: "Hello, mark, anne!",
			items:  []string{"mark", "anne"},
		},
	}

	for _, st := range subtests {
		if s := SayHello(st.items); s != st.result {
			t.Errorf("wanted %s (%v), got %s", st.result, st.items, s)
		}
	}
}
