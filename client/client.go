package client

import (
	"net/http"
)

type Client interface {
	Get(path string, header http.Header, body []byte, result interface{}) (int, error)

	Post(path string, header http.Header, body []byte, result interface{}) (int, error)

	Put(path string, header http.Header, body []byte, result interface{}) (int, error)

	Patch(path string, header http.Header, body []byte, result interface{}) (int, error)

	Delete(path string, header http.Header, body []byte, result interface{}) (int, error)

	Head(path string, header http.Header, body []byte, result interface{}) (int, error)
}
