package aesjson

import (
	"encoding/json"
	"io/ioutil"
)

type aesJson struct{
	// TODO - don't expose Keys
	Keys map[string]string `json:"Keys"`
}

func NewAesJson() *aesJson {
	return &aesJson{
		Keys: make(map[string]string),
	}
}

func (s aesJson) Read(filename string) (map[string]string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &s)
	if err != nil {
		return nil, err
	}
	return s.Keys, nil
}

func (s aesJson) Write(filename string) error {
	data, err := json.Marshal(s)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, data, 0600)
	if err != nil {
		return err
	}
	return nil
}

func (s aesJson) Add(name, url string) error {
	s.Keys[name] = url
	return nil
}