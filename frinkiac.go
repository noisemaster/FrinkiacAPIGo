package frinkiac

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/mitchellh/go-wordwrap"
)

//Frames Holds data about a Frinkiac (or Morbotron) search
type Frames struct {
	ID        int    `json:"Id"`
	Episode   string `json:"Episode"`
	Timestamp int    `json:"Timestamp"`
}

//Episode Holds information about an Episode
type Episode struct {
	Frame     Frames `json:"Frame"`
	Subtitles []struct {
		ID             int    `json:"Id"`
		RepTimestamp   int    `json:"RepresentativeTimestamp"`
		Episode        string `json:"Episode"`
		StartTimestamp int    `json:"StartTimestamp"`
		EndTimestamp   int    `json:"EndTimestamp"`
		Content        string `json:"Content"`
	} `json:"Subtitles"`
	Nearby []Frames `json:"Nearby"`
}

func getFrinkiacFrameData(query string) ([]Frames, error) {
	var info []Frames
	client := &http.Client{}
	r := strings.NewReplacer(" ", "%20")
	req, err := http.NewRequest("GET", "https://frinkiac.com/api/search?q="+r.Replace(query), nil)
	if err != nil {
		return info, err
	}
	req.Header.Set("User-Agent", "Frinkiac_Api_Go/0.1")
	resp, err := client.Do(req)
	if err != nil {
		return info, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return info, err
	}
	if string(body) == "[]" {
		return info, errors.New("No results found for this search")
	}
	json.Unmarshal(body, &info)
	return info, nil
}

func getFrinkiacEpisodeInfo(frame Frames) (Episode, error) {
	var info Episode
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://frinkiac.com/api/caption?e="+frame.Episode+"&t="+strconv.Itoa(frame.Timestamp), nil)
	if err != nil {
		return info, err
	}
	req.Header.Set("User-Agent", "Frinkiac_Api_Go/0.1")
	resp, err := client.Do(req)
	if err != nil {
		return info, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return info, err
	}
	json.Unmarshal(body, &info)
	return info, nil
}

func getMorbotronFrameData(query string) ([]Frames, error) {
	var info []Frames
	client := &http.Client{}
	r := strings.NewReplacer(" ", "%20")
	req, err := http.NewRequest("GET", "https://morbotron.com/api/search?q="+r.Replace(query), nil)
	if err != nil {
		return info, err
	}
	req.Header.Set("User-Agent", "Frinkiac_Api_Go/0.1")
	resp, err := client.Do(req)
	if err != nil {
		return info, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return info, err
	}
	if string(body) == "[]" {
		return info, errors.New("No results found for this search")
	}
	json.Unmarshal(body, &info)
	return info, nil
}

func getMorbotronEpisodeInfo(frame Frames) (Episode, error) {
	var info Episode
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://morbotron.com/api/caption?e="+frame.Episode+"&t="+strconv.Itoa(frame.Timestamp), nil)
	if err != nil {
		return info, err
	}
	req.Header.Set("User-Agent", "Frinkiac_Api_Go/0.1")
	resp, err := client.Do(req)
	if err != nil {
		return info, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return info, err
	}
	json.Unmarshal(body, &info)
	return info, nil
}

//GetFrinkiacFrame Sends a URL of a frame from Frinkiac
func GetFrinkiacFrame(query string) (string, error) {
	frames, err := getFrinkiacFrameData(query)
	if err != nil {
		return "", err
	}
	return "https://frinkiac.com/img/" + frames[0].Episode + "/" + strconv.Itoa(frames[0].Timestamp) + ".jpg", nil
}

//GetFrinkiacMeme Returns a URL of a frame with a caption
func GetFrinkiacMeme(query string) (string, error) {
	frames, err := getFrinkiacFrameData(query)
	if err != nil {
		return "", err
	}
	info, err := getFrinkiacEpisodeInfo(frames[0])
	if err != nil {
		return "", err
	}
	cap := wordwrap.WrapString(info.Subtitles[0].Content, 25)
	uEnc := base64.URLEncoding.EncodeToString([]byte(cap))
	return "https://frinkiac.com/meme/" + frames[0].Episode + "/" + strconv.Itoa(frames[0].Timestamp) + ".jpg?b64lines=" + uEnc, nil
}

//GetFrinkiacGifMeme Returns a URL for a GIF with a caption
func GetFrinkiacGifMeme(query string) (string, error) {
	var cap string
	frames, err := getFrinkiacFrameData(query)
	if err != nil {
		return "", err
	}
	info, err := getFrinkiacEpisodeInfo(frames[0])
	if err != nil {
		return "", err
	}
	for _, v := range info.Subtitles {
		cap += v.Content + "\n"
	}
	cap = strings.TrimSuffix(cap, "\n")
	cap = wordwrap.WrapString(cap, 25)
	uEnc := base64.URLEncoding.EncodeToString([]byte(cap))
	return "https://frinkiac.com/gif/" + info.Frame.Episode + "/" + strconv.Itoa(info.Subtitles[0].StartTimestamp) + "/" + strconv.Itoa(info.Subtitles[len(info.Subtitles)-1].EndTimestamp) + ".gif?b64lines=" + uEnc, nil
}

//GetMorbotronFrame Sends a URL of a frame from Morbotron
func GetMorbotronFrame(query string) (string, error) {
	frames, err := getMorbotronFrameData(query)
	if err != nil {
		return "", err
	}
	return "https://morbotron.com/img/" + frames[0].Episode + "/" + strconv.Itoa(frames[0].Timestamp) + ".jpg", nil
}

//GetMorbotronMeme Returns a URL of a frame with a caption
func GetMorbotronMeme(query string) (string, error) {
	frames, err := getMorbotronFrameData(query)
	if err != nil {
		return "", err
	}
	info, err := getMorbotronEpisodeInfo(frames[0])
	if err != nil {
		return "", err
	}
	cap := wordwrap.WrapString(info.Subtitles[0].Content, 25)
	uEnc := base64.URLEncoding.EncodeToString([]byte(cap))
	return "https://morbotron.com/meme/" + frames[0].Episode + "/" + strconv.Itoa(frames[0].Timestamp) + ".jpg?b64lines=" + uEnc, nil
}

//GetMorbotronGifMeme Returns a URL for a GIF with a caption
func GetMorbotronGifMeme(query string) (string, error) {
	var cap string
	frames, err := getMorbotronFrameData(query)
	if err != nil {
		return "", err
	}
	info, err := getMorbotronEpisodeInfo(frames[0])
	if err != nil {
		return "", err
	}
	for _, v := range info.Subtitles {
		cap += v.Content + "\n"
	}
	cap = strings.TrimSuffix(cap, "\n")
	cap = wordwrap.WrapString(cap, 25)
	uEnc := base64.URLEncoding.EncodeToString([]byte(cap))
	return "https://morbotron.com/gif/" + info.Frame.Episode + "/" + strconv.Itoa(info.Subtitles[0].StartTimestamp) + "/" + strconv.Itoa(info.Subtitles[len(info.Subtitles)-1].EndTimestamp) + ".gif?b64lines=" + uEnc, nil
}
