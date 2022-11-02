package integers

import "testing"

func Test_Add(t *testing.T) {
	desired := 4
	actual := Add(2, 2)

	if actual != desired {
		t.Errorf("expected '%d' but got '%d'", desired, actual)
	}
}
