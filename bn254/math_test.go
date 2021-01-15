package bn254

import "testing"

func Test_bigIntToBE(t *testing.T) {
	if x, y := bigIntToGo(p).Text(16), p.ToString(); x != y {
		t.Fatalf("%s != %s", x, y)
	}
}
