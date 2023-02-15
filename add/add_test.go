package add

import "testing"

func Test_Add(t *testing.T) {
	want := 20
	a := Add(10, 10)

	if want != a {
		t.Logf("want %d but add result is %d", want, a)
	}
}
