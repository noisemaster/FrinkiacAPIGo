package frinkiac

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
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
	json.Unmarshal(body, &info)
	return info, nil
}

func getFrinkiacEpisodeInfo(frame Frames) (Episode, error) {
	var epi Episode

	return epi, nil
}

//GetFrinkiacFrame Sends a URL of a frame from Frinkiac
func GetFrinkiacFrame(query string) (string, error) {
	frames, err := getFrinkiacFrameData(query)
	if err != nil {
		return "", err
	}
	return "https://frinkiac.com/img/" + frames[0].Episode + "/" + string(frames[0].Timestamp) + ".jpg", nil
}
