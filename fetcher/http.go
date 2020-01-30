// Copyright (c) 2020 E-Tiger Studio. All rights reserved.

package fetcher

import (
	"io/ioutil"
	"net/http"
	"subr"
)

type HTTP struct {
	url string
	lastData []byte
}

func (f *HTTP) Fetch(c *subr.Context) error {
	response, err := http.Get(f.url)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	c.Data = contents
	subr.Infoln("HTTP Fetch succeeded")

	return nil
}

func NewHTTP(url string) *HTTP {
	f := &HTTP{
		url:       url,
	}
	return f
}