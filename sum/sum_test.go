package sum

import "testing"

func Test_Sum(t *testing.T) {
	want := 20
	s := Sum(4, 5, 11)

	if s != want {
		t.Logf("want %d but go %d", want, s)
	}
}
