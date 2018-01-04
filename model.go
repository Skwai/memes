package main

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
  "sync"
)

const APIBaseURL = "https://api.giphy.com"
const APIKey = "y982M4dvoKqweMEoRbPbWMqz2agb1BcW"

type Meme struct {
	Id		string
	Url		string		`json:"image_mp4_url"`
}

func (m *Meme) getMeme(id string) error {
  g, err := giphyRequest("/" + id)

  if (err != nil) {
    return err
  }

  m.Id = g.Data.Id
  m.Url = g.Data.Url

  return nil
}

func getRandomMemes(count int) ([]Meme, error) {
	var m []Meme
	var wg sync.WaitGroup

	i := 0

	wg.Add(count)

	for i < count {
		go func() {
			defer wg.Done()
			resp, err := getRandomMeme()

			if err == nil {
				m = append(m, resp)
			}
		}()
		i++
	}

	wg.Wait()

	return m, nil
}

type Gif struct {
	Data Meme
}

func getRandomMeme() (Meme, error) {
	var m Meme

	g, err := giphyRequest("random")

	if err != nil {
		return m, err
	}

	if err != nil {
		return m, err
	}

	m = g.Data

	return m, nil
}

func giphyRequest(path string) (Gif, error) {
  var g Gif
  url := APIBaseURL + "/v1/gifs" + path + "?api_key=" + APIKey
  resp, err := http.Get(url)

	if err != nil {
		return g, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return g, err
  }

  err = json.Unmarshal(body, &g)

  if err != nil {
    return g, err
  }

	return g, nil
}
