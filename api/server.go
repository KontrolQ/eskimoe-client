// This package will contain all the Endpoints for the Eskimoe Server API.

package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type ServerInfo struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Version string `json:"version"`
}

func GetServerInfo(URL string) (ServerInfo, error) {
	_, err := url.ParseRequestURI(strings.TrimSpace(URL))
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

func JoinServerAsMember(serverURL string, member JoinMemberRequest) (JoinMemberSuccessResponse, error) {
	endpoint := serverURL + "/join"

	memberJSON, err := json.Marshal(member)
	if err != nil {
		return JoinMemberSuccessResponse{}, errors.New("failed to encode member data")
	}

	response, err := http.Post(endpoint, "application/json", bytes.NewBuffer(memberJSON))

	if err != nil {
		return JoinMemberSuccessResponse{}, errors.New("failed to send join request")
	}

	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)

	if err != nil {
		return JoinMemberSuccessResponse{}, errors.New("failed to read request response")
	}

	if response.StatusCode != http.StatusCreated {
		var errorResponse JoinMemberErrorResponse
		err = json.Unmarshal(responseData, &errorResponse)

		if err != nil {
			return JoinMemberSuccessResponse{}, errors.New("failed to decode error response")
		}

		return JoinMemberSuccessResponse{}, errors.New(errorResponse.Error)
	}

	var successResponse JoinMemberSuccessResponse

	err = json.Unmarshal(responseData, &successResponse)

	if err != nil {
		return JoinMemberSuccessResponse{}, errors.New("failed to decode success response")
	}

	return successResponse, nil
}
