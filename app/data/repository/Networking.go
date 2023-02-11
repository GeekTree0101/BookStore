package repository

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Networking struct {
	client *http.Client
}

func DefaultNetworking() Networking {
	return Networking{
		client: &http.Client{},
	}
}

func (networking *Networking) fetch(req *http.Request, value any) error {
	res, err := networking.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	bytes, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(bytes, value)
	if err != nil {
		return err
	}
	return nil
}
