package drudge_test

import (
	"testing"

	"github.com/DeedleFake/drudge"
)

func TestClient(t *testing.T) {
	// TODO: Better tests. The bane of small projects...

	var c drudge.Client

	top, err := c.Top()
	if err != nil {
		t.Fatal(err)
	}
	for _, article := range top {
		t.Logf("%#v", article)
	}

	for i := 1; i <= 3; i++ {
		col, err := c.Column(drudge.Column(i))
		if err != nil {
			t.Fatal(err)
		}
		for _, article := range col {
			t.Logf("%#v", article)
		}
	}
}
