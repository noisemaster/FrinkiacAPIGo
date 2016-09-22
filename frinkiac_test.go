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
	final := "https://frinkiac.com/meme/S07E21/579294.jpg?b64lines=T2gsIG5vdCBpbiBVdGljYSwgbm8sCml0J3MgYW4gQWxiYW55CmV4cHJlc3Npb24uCg=="
	recv, err := GetFrinkiacMeme("Steamed Hams")
	if err != nil {
		t.Error(err)
	}
	if recv != final {
		t.Errorf("Recv: %s \n Want: %s \n %s\n", recv, final, "URLs were not equal")
	}
}

func TestGifMeme(t *testing.T) {
	final := "https://frinkiac.com/gif/S07E21/576200/581967.gif?b64lines=T2gsIG5vdCBpbiBVdGljYSwgbm8sCml0J3MgYW4gQWxiYW55CmV4cHJlc3Npb24u"
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
	final := "https://morbotron.com/meme/S01E05/1112308.jpg?b64lines=V2UgYXJlIHRoZSB0cnVlIHJ1bGVycyBvZgp0aGlzIHBsYW5ldC4KSGFuZC1jYXJ2ZWQgZnJvbQptZXRlb3JpdGVzIGJ5IHRoZSByb2JvdApmb3VuZGVycwo="
	recv, err := GetMorbotronMeme("this planet")
	if err != nil {
		t.Error(err)
	}
	if recv != final {
		t.Errorf("Recv: %s \n Want: %s \n %s\n", recv, final, "URLs were not equal")
	}
}

func TestMorboGifMeme(t *testing.T) {
	final := "https://morbotron.com/gif/S01E05/1109389/1115063.gif?b64lines=VGhhdCB3YXMganVzdCBhIHNob3cgZm9yCnRoZSBwdWJsaWMuCldlIGFyZSB0aGUgdHJ1ZSBydWxlcnMgb2YKdGhpcyBwbGFuZXQuCkhhbmQtY2FydmVkIGZyb20KbWV0ZW9yaXRlcyBieSB0aGUgcm9ib3QKZm91bmRlcnM="
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
