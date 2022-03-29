package gowebshell

import (
	"crypto/tls"
	"net/http"
)

var DefaultOption *Options


func init()  {
	tr := &http.Transport{
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
	}
	DefaultOption=&Options{}
	DefaultOption.client=&http.Client{Transport: tr}
	DefaultOption.randomlen=10
}
