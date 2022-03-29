package gowebshell

import "net/http"

type Options struct {
	randomlen	int
	client	*http.Client
}
