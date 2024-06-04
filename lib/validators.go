package lib

import (
	"errors"
	"eskimoe-client/api"
)

func ServerReachableValidator() func(string) error {
	return func(s string) error {
		serverInfo, err := api.GetServerInfo(s)
		if err != nil {
			return err
		}
		if serverInfo.Name == "" {
			return errors.New("server responded with an empty name")
		}
		return nil
	}
}
