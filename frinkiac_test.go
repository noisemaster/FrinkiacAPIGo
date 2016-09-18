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

func TestMeme(t *testing.T) {
	final := "https://frinkiac.com/meme/S07E21/579294.jpg?b64lines=YW5kIEkndmUgbmV2ZXIgaGVhcmQKYW55b25lIHVzZSB0aGUgcGhyYXNlCiJzdGVhbWVkIGhhbXMuIg=="
	recv, err := GetFrinkiacFrameAndCaption("Steamed Hams")
	if err != nil {
		t.Error(err)
	}
	if recv != final {
		t.Errorf("Recv: %s \n Want: %s \n %s\n", recv, final, "URLs were not equal")
	}
}
