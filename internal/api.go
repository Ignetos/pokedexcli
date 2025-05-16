package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

const BASEURL = "https://pokeapi.co/api/v2/"

type MapData struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func errIfNotHTTPS(URL string) error {
	url, err := url.Parse(URL)
	if err != nil {
		return err
	}
	if url.Scheme != "https" {
		return fmt.Errorf("URL scheme is not HTTPS: %s", URL)
	}
	return nil
}

func GetMapData(url string) (MapData, error) {
	if err := errIfNotHTTPS(url); err != nil {
		return MapData{}, err
	}

	res, err := http.Get(url)
	if err != nil {
		return MapData{}, errors.New("error getting response")
	}
	defer res.Body.Close()

	var maps MapData
	err = json.NewDecoder(res.Body).Decode(&maps)
	if err != nil {
		return MapData{}, errors.New("error decoding response")
	}

	return maps, nil
}
