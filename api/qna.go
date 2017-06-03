package api

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

func GetAnswer(enpoint, question string) (string, error) {
	resp, err := http.PostForm(enpoint, url.Values{"question": {question}})
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return string(body), nil
}
