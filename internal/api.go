package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

type ExploreData struct {
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
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

func UnmarshalMapData(b []byte) (MapData, error) {
	var data MapData
	err := json.Unmarshal(b, &data)
	if err != nil {
		return MapData{}, fmt.Errorf("error decoding response: %w", err)
	}
	return data, nil
}

func UnmarshalExploreData(b []byte) (ExploreData, error) {
	var data ExploreData
	err := json.Unmarshal(b, &data)
	if err != nil {
		return ExploreData{}, fmt.Errorf("error decoding response: %w", err)
	}
	return data, nil
}

func GetMapData(url string, cache *Cache) (MapData, error) {
	if err := errIfNotHTTPS(url); err != nil {
		return MapData{}, err
	}

	cachedData, ok := cache.Get(url)
	if ok {
		return UnmarshalMapData(cachedData)
	}

	res, err := http.Get(url)
	if err != nil {
		return MapData{}, errors.New("error getting response")
	}
	defer res.Body.Close()

	byteData, err := io.ReadAll(res.Body)
	if err != nil {
		return MapData{}, fmt.Errorf("error reading response: %w", err)
	}
	cache.Add(url, byteData)

	cachedData, _ = cache.Get(url)
	return UnmarshalMapData(cachedData)
}

func GetExploreData(url string, cache *Cache) (ExploreData, error) {
	if err := errIfNotHTTPS(url); err != nil {
		return ExploreData{}, err
	}

	cachedData, ok := cache.Get(url)
	if ok {
		return UnmarshalExploreData(cachedData)
	}

	res, err := http.Get(url)
	if err != nil {
		return ExploreData{}, errors.New("error getting response")
	}
	if res.StatusCode > 299 {
		return ExploreData{}, nil
	}
	defer res.Body.Close()

	byteData, err := io.ReadAll(res.Body)
	if err != nil {
		return ExploreData{}, fmt.Errorf("error reading response: %w", err)
	}
	cache.Add(url, byteData)

	cachedData, _ = cache.Get(url)
	return UnmarshalExploreData(cachedData)
}
