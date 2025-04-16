package internal

import (
	"io"
	"net/http"
)

func GetCachedData(c *Config, url string) ([]byte, error) {
	var data []byte
	if c.Cache != nil {
		if cachedData, exists := c.Cache.Get(url); exists {
			data = cachedData
		}
	}
	if len(data) == 0 {
		locs, err := getDataFromApi(url)
		data = locs
		c.Cache.Add(url, data)
		if err != nil {
			return nil, err
		}
	}
	return data, nil
}

func getDataFromApi(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	return data, nil

}
