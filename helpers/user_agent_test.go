package helpers

import (
	"testing"
)

func TestIsBrowser(t *testing.T) {
	ua := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:126.0) Gecko/20100101 Firefox/126.0"

	want := true
	got := IsBrowser(ua)

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}
