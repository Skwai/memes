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

	url := APIBaseURL + "/v1/gifs/random?api_key=" + APIKey	
	resp, err := http.Get(url)

	if err != nil {
		return m, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return m, err
	}

	var g Gif

	json.Unmarshal(body, &g)	

	if err != nil {
		return m, err
	}

	m = g.Data

	return m, nil
}