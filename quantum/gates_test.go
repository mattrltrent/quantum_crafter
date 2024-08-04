package quantum

import "testing"

// dummy test to basically see the tests suite working
func TestGatesIdentity(t *testing.T) {
	identity := Identity(2)
	if identity.WiresNeeded() != 1 {
		t.Errorf("gate should need 1 wire")
	}
	if identity.Name() != "I" {
		t.Errorf("gate should have name I")
	}
	if identity.Data().Rows != 2 || identity.Data().Cols != 2 {
		t.Errorf("gate should have 2x2 matrix")
	}
	if identity.Data().Data[0][0] != 1 || identity.Data().Data[1][1] != 1 {
		t.Errorf("gate should have 1s on the diagonal")
	}
}
