package bn254

import "testing"

func Test_bigIntFromGo(t *testing.T) {
	pMinusN := p.Minus(n)
	expect := pMinusN.ToString()

	if got := bigIntFromGo(bigIntToGo(pMinusN)).ToString(); expect != got {
		t.Fatalf("\nexpect: %s,\n   got: %s", expect, got)
	}
}

func Test_bigIntToGo(t *testing.T) {
	if x, y := bigIntToGo(p).Text(16), p.ToString(); x != y {
		t.Fatalf("%s != %s", x, y)
	}
}
