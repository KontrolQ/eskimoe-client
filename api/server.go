// This package will contain all the Endpoints for the Eskimoe Server API.

package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func GetServerInfo(URL string) (ServerInfoUnauthorized, error) {
	_, err := url.ParseRequestURI(strings.TrimSpace(URL))
	if err != nil {
		return ServerInfoUnauthorized{}, errors.New("provided URL is not valid")
	}

	response, err := http.Get(URL)
	if err != nil {
		return ServerInfoUnauthorized{}, errors.New("server is unreachable")
	}

	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return ServerInfoUnauthorized{}, errors.New("server did not return a valid response")
	}

	var serverInfo ServerInfoUnauthorized

	err = json.Unmarshal(responseData, &serverInfo)
	if err != nil {
		return ServerInfoUnauthorized{}, errors.New("server is not an eskimoe server")
	}

	return serverInfo, nil
}

func GetAuthorizedServerInfo(URL string, authToken string) (ServerInfoAuthorized, error) {
	_, err := url.ParseRequestURI(strings.TrimSpace(URL))
	if err != nil {
		return ServerInfoAuthorized{}, errors.New("provided URL is not valid")
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return ServerInfoAuthorized{}, errors.New("failed to create request")
	}

	req.Header.Set("Authorization", authToken)

	response, err := client.Do(req)
	if err != nil {
		return ServerInfoAuthorized{}, errors.New("server is unreachable")
	}

	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return ServerInfoAuthorized{}, errors.New("server did not return a valid response")
	}

	var serverInfo ServerInfoAuthorized

	err = json.Unmarshal(responseData, &serverInfo)
	if err != nil {
		return ServerInfoAuthorized{}, errors.New("server is not an eskimoe server")
	}

	return serverInfo, nil
}

func JoinServerAsMember(serverURL string, member JoinMemberRequest) (JoinMemberSuccessResponse, error) {
	serverURL = strings.TrimRight(serverURL, "/")

	endpoint := serverURL + "/members/join"

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

func LeaveServer(serverURL string, authToken string) error {
	serverURL = strings.TrimRight(serverURL, "/")

	endpoint := serverURL + "/members/leave"

	client := &http.Client{}

	req, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		return errors.New("failed to create request")
	}

	req.Header.Set("Authorization", authToken)

	response, err := client.Do(req)
	if err != nil {
		return errors.New("failed to send request")
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return errors.New("failed to leave server")
	}

	return nil
}

func GetMessagesInRoom(serverURL string, roomId int, authToken string) ([]Message, error) {
	serverURL = strings.TrimRight(serverURL, "/")
	endpoint := fmt.Sprintf("%s/rooms/%d/messages", serverURL, roomId)

	client := &http.Client{}

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, errors.New("failed to create request")
	}

	req.Header.Set("Authorization", authToken)

	response, err := client.Do(req)
	if err != nil {
		return nil, errors.New("failed to send request")
	}

	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("failed to read response")
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("failed to get messages")
	}

	var messages []Message

	err = json.Unmarshal(responseData, &messages)
	if err != nil {
		return nil, errors.New("failed to decode messages")
	}

	return messages, nil
}

func SendMessageToRoom(serverURL string, roomId int, authToken string, message SendRoomMessage) error {
	serverURL = strings.TrimRight(serverURL, "/")
	endpoint := fmt.Sprintf("%s/rooms/%d/messages/new", serverURL, roomId)

	messageJSON, err := json.Marshal(message)

	if err != nil {
		return errors.New("failed to encode message")
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(messageJSON))
	if err != nil {
		return errors.New("failed to create request")
	}

	req.Header.Set("Authorization", authToken)
	req.Header.Set("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		return errors.New("failed to send request")
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusCreated {
		fmt.Println(response)
		return errors.New("failed to send message")
	}

	return nil
}

func DeleteMessageInRoom(serverURL string, roomId int, messageId int, authToken string) error {
	serverURL = strings.TrimRight(serverURL, "/")
	endpoint := fmt.Sprintf("%s/rooms/%d/messages/%d", serverURL, roomId, messageId)

	client := &http.Client{}

	req, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		return errors.New("failed to create request")
	}

	req.Header.Set("Authorization", authToken)

	response, err := client.Do(req)
	if err != nil {
		return errors.New("failed to send request")
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return errors.New("failed to delete message")
	}

	return nil
}
