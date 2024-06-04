// This package will contain all the Endpoints for the Eskimoe Server API.

package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

type ServerInfo struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Version string `json:"version"`
}

func GetServerInfo(URL string) (ServerInfo, error) {
	_, err := url.ParseRequestURI(URL)
	if err != nil {
		return ServerInfo{}, errors.New("provided URL is not valid")
	}

	response, err := http.Get(URL)
	if err != nil {
		return ServerInfo{}, errors.New("server is unreachable")
	}

	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return ServerInfo{}, errors.New("server did not return a valid response")
	}

	var serverInfo ServerInfo

	err = json.Unmarshal(responseData, &serverInfo)
	if err != nil {
		return ServerInfo{}, errors.New("server is not an eskimoe server")
	}

	return serverInfo, nil
}
