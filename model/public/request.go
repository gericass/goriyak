package public

import (
	"net/http"
	"bytes"
	"errors"
)

func HTTPError(res *http.Response) error {
	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	return errors.New(buf.String())
}

// GetRequest : receives the endpoint and returns the result of the request
func GetRequest(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func PutRequest(url string, body string) (*http.Response, error) {
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func DeleteRequest(url string) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
