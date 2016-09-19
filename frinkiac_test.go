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
	recv, err := GetFrinkiacMeme("Steamed Hams")
	if err != nil {
		t.Error(err)
	}
	if recv != final {
		t.Errorf("Recv: %s \n Want: %s \n %s\n", recv, final, "URLs were not equal")
	}
}

func TestGifMeme(t *testing.T) {
	final := "https://frinkiac.com/gif/S07E21/576200/581967.gif?b64lines=YW5kIEkndmUgbmV2ZXIgaGVhcmQKYW55b25lIHVzZSB0aGUgcGhyYXNlCiJzdGVhbWVkIGhhbXMuIgpPaCwgbm90IGluIFV0aWNhLCBubywKaXQncyBhbiBBbGJhbnkKZXhwcmVzc2lvbi4="
	recv, err := GetFrinkiacGifMeme("Steamed Hams")
	if err != nil {
		t.Error(err)
	}
	if recv != final {
		t.Errorf("Recv: %s \n Want: %s \n %s\n", recv, final, "URLs were not equal")
	}
}

func TestMorboFrame(t *testing.T) {
	final := "https://morbotron.com/img/S01E05/1112308.jpg"
	recv, err := GetMorbotronFrame("this planet")
	if err != nil {
		t.Error(err)
	}
	if recv != final {
		t.Errorf("Recv: %s \n Want: %s \n %s\n", recv, final, "URLs were not equal")
	}
}

func TestMorboMeme(t *testing.T) {
	final := "https://morbotron.com/meme/S01E05/1112308.jpg?b64lines=VGhhdCB3YXMganVzdCBhIHNob3cgZm9yCnRoZSBwdWJsaWMu"
	recv, err := GetMorbotronMeme("this planet")
	if err != nil {
		t.Error(err)
	}
	if recv != final {
		t.Errorf("Recv: %s \n Want: %s \n %s\n", recv, final, "URLs were not equal")
	}
}

func TestMorboGifMeme(t *testing.T) {
	final := "https://morbotron.com/gif/S01E05/1108431/1115063.gif?b64lines=VGhhdCB3YXMganVzdCBhIHNob3cgZm9yCnRoZSBwdWJsaWMuCldlIGFyZSB0aGUgdHJ1ZSBydWxlcnMgb2YKdGhpcyBwbGFuZXQuCkhhbmQtY2FydmVkIGZyb20KbWV0ZW9yaXRlcyBieSB0aGUgcm9ib3QKZm91bmRlcnM="
	recv, err := GetMorbotronGifMeme("this planet")
	if err != nil {
		t.Error(err)
	}
	if recv != final {
		t.Errorf("Recv: %s \n Want: %s \n %s\n", recv, final, "URLs were not equal")
	}
}

func TestGarbage(t *testing.T) {
	_, err := GetMorbotronGifMeme("soafgnme")
	if err == nil {
		t.Error(err)
	}
}
