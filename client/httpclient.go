package client

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"log"
)

type HttpClient struct {
	url    string
	user   string
	pass   string
	client *http.Client
}

func ClientInstance(url, user, pass string) Client {

	return &HttpClient{
		url:    url,
		user:   user,
		pass:   pass,
		client: &http.Client{Timeout: 10 * time.Second},
	}

}

func (c *HttpClient) Post(path string, header http.Header, body []byte, result interface{}) (int, error) {

	return c.getParsedResponse("POST", path, header, bytes.NewReader(body), result)

}

func (c *HttpClient) Get(path string, header http.Header, body []byte, result interface{}) (int, error) {

	return c.getParsedResponse("GET", path, header, bytes.NewReader(body), result)

}

func (c *HttpClient) Put(path string, header http.Header, body []byte, result interface{}) (int, error) {

	return c.getParsedResponse("PUT", path, header, bytes.NewReader(body), result)

}

func (c *HttpClient) Patch(path string, header http.Header, body []byte, result interface{}) (int, error) {

	return c.getParsedResponse("PATCH", path, header, bytes.NewReader(body), result)

}

func (c *HttpClient) Delete(path string, header http.Header, body []byte, result interface{}) (int, error) {

	return c.getParsedResponse("DELETE", path, header, bytes.NewReader(body), result)

}

func (c *HttpClient) Head(path string, header http.Header, body []byte, result interface{}) (int, error) {

	return c.getParsedResponse("HEAD", path, header, bytes.NewReader(body), result)

}

//BasicAuthEncode Authorization
func basicAuthEncode(user, pass string) string {
	if user != "" && pass != "" {
		return base64.StdEncoding.EncodeToString([]byte(user + ":" + pass))
	}
	return ""
}

func (c *HttpClient) doRequest(method, path string, header http.Header, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, c.url+path, body)
	log.Println("request :", method, c.url+path)
	if err != nil {
		return nil, err
	}

	if header != nil {
		for k, v := range header {
			req.Header[k] = v
		}
	}

	if baseAuth := basicAuthEncode(c.user, c.pass); baseAuth != "" {
		req.Header.Set("Authorization", "Basic "+baseAuth)
	}

	for k, v := range header {
		req.Header[k] = v
	}

	return c.client.Do(req)
}

func (c *HttpClient) getResponse(method, path string, header http.Header, body io.Reader) ([]byte, int, error) {
	resp, err := c.doRequest(method, path, header, body)
	if err != nil {
		return nil, 500, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 500, err
	}

	switch resp.StatusCode {
	case 400:
		return nil, 400, errors.New("400 Unsatisfied Constraints")
	case 403:
		return nil, 403, errors.New("403 Forbidden")
	case 404:
		return nil, 404, errors.New("404 Not Found")
	case 409:
		return nil, 409, errors.New("409 Already Exists")
	case 500:
		return nil, 500, errors.New("500 Internal Errors")
	}
	if location := resp.Header.Get("location"); location != "" {
		location = fmt.Sprintf(`"%s"`, location)
		data = []byte(`{"location":` + location + `}`)
	}

	if resp.StatusCode/100 != 2 {
		errMap := make(map[string]interface{})
		if err = json.Unmarshal(data, &errMap); err != nil {
			return nil, resp.StatusCode, err
		}
		return nil, resp.StatusCode, errors.New(errMap["message"].(string))
	}

	return data, resp.StatusCode, nil
}

func (c *HttpClient) getParsedResponse(method, path string, header http.Header, body io.Reader, obj interface{}) (int, error) {
	data, status, err := c.getResponse(method, path, header, body)
	if err != nil {
		return status, err
	}

	if obj == nil {
		return status, err
	}
	return status, json.Unmarshal(data, obj)
}
