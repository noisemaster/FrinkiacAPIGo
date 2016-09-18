package frinkiac

import "testing"

func TestFrame(t *testing.T) {
	final := "https://frinkiac.com/img/S07E21/579294.jpg"
	recv, err := GetFrinkiacFrame("Steamed Hams")
	if err != nil {
		t.Error(err)
	}
	if recv != final {
		t.Errorf("Recv: %s \n Want: %s \n %s\n", recv, final, "URLs were not equal")
	}
}
