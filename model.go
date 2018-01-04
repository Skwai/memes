package main

// https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql

import (
  "net/http"
	"encoding/json"
	"io/ioutil"
  "sync"
)

const APIBaseURL = "https://api.giphy.com/v1/gifs"
const APIKey = "y982M4dvoKqweMEoRbPbWMqz2agb1BcW"

type GiphyResponse struct {
  Data Gif `json:"data"`
  Meta struct {
    Msg             string
    ResponseID      string
    Status          int
  }
}

type Gif struct {
  ID                string    `json:"id"`
  Type              string    `json:"type"`
  URL               string    `json:"url"`
  ImageOriginalURL  string    `json:"image_original_url"`
  ImageMp4Url       string    `json:"image_mp4_url"`
  ImageUrl          string    `json:"image_url"`
  ImageWidth        string    `json:"image_width"`
  ImageHeight       string    `json:"image_height"`
}

func (m *Gif) getMeme(id string) error {
  /*
  g, err := giphyRequest("/" + id)

  if (err != nil) {
    return err
  }
  */
  return nil
}

func getRandomMemes(count int) ([]Gif, error) {
	var g []Gif
	var wg sync.WaitGroup

	i := 0

	wg.Add(count)

	for i < count {
		go func() {
			defer wg.Done()
			resp, err := getRandomMeme()

			if err == nil {
				g = append(g, resp)
			}
		}()
		i++
	}

	wg.Wait()

	return g, nil
}

func getRandomMeme() (Gif, error) {
  var g GiphyResponse

	g, err := giphyRequest("/random")

	if err != nil {
		return g.Data, err
  }

	if err != nil {
		return g.Data, err
	}

	return g.Data, nil
}

func giphyRequest(path string) (GiphyResponse, error) {
  var g GiphyResponse
  url := APIBaseURL + path + "?api_key=" + APIKey
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
